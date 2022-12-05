package configs

import (
	"cas/tools"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	DBConfig `mapstructure:"db"`

	Code `mapstructure:"code"`

	APIConfig `mapstructure:"api"`
}

type Code struct {
	Invite string
}

type APIConfig struct {
	HttpPort int `mapstructure:"http_port"`
	GrpcPort int `mapstructure:"grpc_port"`
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func InitConfig() (err error) {
	// 默认配置文件路径
	configPath := tools.GetRootPath("/config.yaml")
	logrus.Printf("===> config path: %s", configPath)
	// 初始化配置文件
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	// 观察配置文件变动
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Printf("config file has changed")
		if err = viper.Unmarshal(&Conf); err != nil {
			logrus.Errorf("failed at unmarshal config file after change, err: %v", err)
		}
	})
	// 将配置文件读入 viper
	if err = viper.ReadInConfig(); err != nil {
		logrus.Errorf("failed at ReadInConfig, err: %v", err)
	}
	// 解析到变量中
	if err = viper.Unmarshal(&Conf); err != nil {
		logrus.Errorf("failed at Unmarshal config file, err: %v", err)
	}
	// 从环境变量中覆盖配置
	//viper.AutomaticEnv()
	// 返回 nil 或错误
	return err
}
