//go:build ignore
// +build ignore

package main

import (
	"ariga.io/atlas/sql/sqltool"
	"cas/configs"
	"cas/pkg/ent/migrate"
	"cas/tools"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func main() {
	// 创建文件夹
	err := os.Mkdir("../internal/db/migrations", 0777)
	if err != nil && !os.IsExist(err) {
		logrus.Errorf("failed at mkdir migrations")
	}
	ctx := context.Background()
	// 指定文件夹
	dir, err := sqltool.NewGolangMigrateDir("../internal/db/migrations")
	if err != nil {
		logrus.Errorf("failed at creating atlas migration directory: %v", err)
		return
	}
	// 需要初始化数据库配置
	configPath := fmt.Sprintf("%s", filepath.Join(tools.GetDeployPath(""), "..", "config.yaml"))
	err = configs.InitConfig(configPath)
	if err != nil {
		return
	}
	dbConf := configs.Conf.DBConfig
	// 迁移条件
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
	dbUrl := fmt.Sprintf("%s://%s:%s@%s:%v/%s?sslmode=disable&TimeZone=Asia/Shanghai", dbConf.Driver, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	// 开始创建迁移文件
	err = migrate.NamedDiff(ctx, dbUrl, "update", opts...)
	if err != nil {
		logrus.Errorf("failed at generating migrations, err: %v", err)
	}
	return
}
