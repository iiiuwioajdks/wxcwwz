package global

import "gorm.io/gorm"

type Server struct {
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	DBName   string `mapstructure:"dbname"`
	Port     int    `mapstructure:"port"`
}

var Conf *Server

var DB *gorm.DB
