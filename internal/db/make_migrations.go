// +build ignore

package main

import (
	"ariga.io/atlas/sql/sqltool"
	"cas/configs"
	"cas/pkg/ent/migrate"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	// 指定文件夹
	dir, err := sqltool.NewGolangMigrateDir("migrations")
	if err != nil {
		logrus.Errorf("failed at creating atlas migration directory: %v", err)
		return
	}
	// 迁移条件
	dbConf := configs.Conf.DBConfig
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		// 配合 golang-migrate 使用
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
		// 步进模式
		schema.WithMigrationMode(schema.ModeInspect),
		// 指定数据库类型
		schema.WithDialect(dbConf.Driver),
		// 移除外键约束
		schema.WithForeignKeys(false),
		// 不可删字段
		schema.WithDropColumn(false),
		// 可删索引
		schema.WithDropIndex(true),
	}
	// 需知道数据库目标
	url = fmt.Sprintf("%s://%s:%s@%s:%v/%s?sslmodel=disable&TimeZone=Asia/Shanghai", dbConf.Driver, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	// 开始创建迁移文件
	err = migrate.NamedDiff(ctx, url, "update", opts...)
	if err != nil {
		logrus.Errorf("failed at generating migrations, err: %v", err)
	}
	return
}