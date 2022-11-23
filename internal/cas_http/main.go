package main

import (
	"cas/configs"
	"cas/internal/db"
	"cas/pkg/ent"
	"cas/pkg/graphql"
	"cas/tools"
	"context"
	"database/sql"
	"database/sql/driver"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	err = configs.InitLogger()
	err = configs.InitConfig("")
	err = tools.Init()
	err = db.InitDB()
	if err != nil {
		return
	}

	//// 初始化 http 服务，使用 gqlgen 的操作台
	//http.Handle("/", playground.Handler("Test", "/graphql"))
	//logrus.Printf("Listening on :%v", configs.Conf.APIConfig.HttpPort)
	//if err = http.ListenAndServe(fmt.Sprintf(":%v", configs.Conf.APIConfig.HttpPort), nil); err != nil {
	//	logrus.Fatalf("http with graphql server DOWN!, err: %v", err)
	//}
	r := gin.Default()
	r.POST("/graph", graphqlHandler())
	r.GET("/", playgroundHandler())
	err = r.Run(fmt.Sprintf(":%v", configs.Conf.APIConfig.HttpPort))
	if err != nil {
		return
	}
}

func graphqlHandler() gin.HandlerFunc {
	// 创建数据库链接
	client := db.NewDBClient()
	// 初始化 graphql server
	srv := handler.NewDefaultServer(graphql.NewSchema(client))
	// 自定义事务隔离等级
	srv.Use(entgql.Transactioner{
		TxOpener: entgql.TxOpenerFunc(func(ctx context.Context) (context.Context, driver.Tx, error) {
			tx, err := client.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
			if err != nil {
				return nil, nil, err
			}
			ctx = ent.NewTxContext(ctx, tx)
			ctx = ent.NewContext(ctx, tx.Client())
			return ctx, tx, nil
		}),
	})
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	srv := playground.Handler("Test", "/graph")
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}
