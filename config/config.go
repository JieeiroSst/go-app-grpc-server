package config 

import (
	"github.com/JIeeiroSst/go-app/repositories/mysql"
	"github.com/JIeeiroSst/go-app/repositories/mongo"
	"github.com/JIeeiroSst/go-app/log"
	"github.com/kelseyhightower/envconfig"
)

type WebConfig struct {
	PORT string     			`envconfig:"WEB_PORT"`
	MysqlConfig mysql.Config 	`envconfig:"WEB_MYSQL"`
	MongoCofig mongo.Mongoconn 	`envconfig:"WEB_MONGO"`
	Nats string 				`envconfig:"WEB_Nats"`
}

var Config WebConfig

func init(){
	err:=envconfig.Process("",&Config)
	if err!=nil{
			panic(err)
	}
	log.InitZapLog().Info(Config.PORT)
}