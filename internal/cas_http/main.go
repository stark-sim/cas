package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/stark-sim/cas/configs"
	httpMiddlewares "github.com/stark-sim/cas/internal/cas_http/middlewares"
	"github.com/stark-sim/cas/internal/db"
	"github.com/stark-sim/cas/pkg/ent"
	"github.com/stark-sim/cas/pkg/graphql"
	"github.com/stark-sim/cas/pkg/graphql/middlewares"
	"github.com/stark-sim/cas/tools"
)

func main() {
	var err error
	err = configs.InitLogger()
	err = configs.InitConfig()
	err = tools.Init()
	err = db.InitDB()
	if err != nil {
		panic(err)
	}
	// 结合 gin 启动 http 服务
	r := gin.Default()
	r.Use(middlewares.WriterMiddleware())
	r.Use(httpMiddlewares.CORS())
	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	err = r.Run(fmt.Sprintf(":%v", configs.Conf.APIConfig.HttpPort))
	if err != nil {
		panic(err)
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
	// 接上 cookie 校验中间件
	//srv.Use(middlewares.NewAuthenticationMiddleware("login", "register"))
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	srv := playground.Handler("Test", "/cas/graphql")
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
