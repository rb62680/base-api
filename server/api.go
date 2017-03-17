package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

type API struct {
	Router *gin.Engine
	Config *viper.Viper
	Database *mgo.Database
}
