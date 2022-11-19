package db

import (
	"cas/configs"
	"errors"
	"fmt"
	golangMigrate "github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
)

var url = ""

func InitDB() (err error) {
	dbConf := configs.Conf.DBConfig
	url = fmt.Sprintf("%s://%s:%s@%s:%v/%s?sslmodel=disable&TimeZone=Asia/Shanghai", dbConf.Driver, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	err = migrateWithMigrationFiles()
	return err
}

func migrateWithMigrationFiles() (err error) {
	m, err := golangMigrate.New("file://migrations", url)
	if err != nil {
		logrus.Errorf("failed at new migrate, err: %v", err)
		return err
	}
	// 没有变更不算错误
	if err = m.Up(); err != nil && !errors.Is(err, golangMigrate.ErrNoChange) {
		logrus.Errorf("faied at migrating: %v", err)
		return err
	}
	return err
}