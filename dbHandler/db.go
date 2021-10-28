package dbHandler

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"wvCheck/global"
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		global.Conf.UserName,
		global.Conf.Password,
		global.Conf.Host,
		global.Conf.Port,
		global.Conf.DBName,
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 不用复数
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func InitConfig() error {
	workdir, _ := os.Getwd()
	fmt.Println(workdir)
	viper.SetConfigFile(workdir + "/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "setting init error")
	}
	if err := viper.Unmarshal(&global.Conf); err != nil {
		return errors.Wrap(err, "unmarshal init error")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return nil
}
