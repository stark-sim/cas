package db

import (
	"cas/configs"
	"cas/pkg/ent"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var url = ""

func InitDB() (err error) {
	dbConf := configs.Conf.DBConfig
	url = fmt.Sprintf("%s://%s:%s@%s:%v/%s?sslmode=disable&TimeZone=Asia/Shanghai", dbConf.Driver, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	err = migrateWithMigrationFiles()
	return err
}

func migrateWithMigrationFiles() (err error) {
	m, err := migrate.New("file://internal/db/migrations", url)
	if err != nil {
		logrus.Errorf("failed at new migrate, err: %v", err)
		return err
	}
	// 没有变更不算错误
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logrus.Errorf("faied at migrating: %v", err)
		return err
	}
	return nil
}

func NewDBClient() *ent.Client {
	dbConf := configs.Conf.DBConfig
	dataSourceName := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Database)
	logrus.Debugf("dsn: %s\n", dataSourceName)
	client, err := ent.Open(dbConf.Driver, dataSourceName)
	if err != nil {
		logrus.Errorf("failed at creating ent client with db %s, err: %v", dataSourceName, err)
		return nil
	}
	return client
}