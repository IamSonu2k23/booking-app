package app

import (
	"github.com/gin-gonic/gin"

	"ticket-booking/api"
	"ticket-booking/config/dbconfig/dbinit"
	"ticket-booking/config/envconfig"
	"ticket-booking/global"
)

var GinApp *gin.Engine

func Init() {
	//initialize env parameters
	envconfig.InitEnvVars()

	//initialize database
	dbinit.Init()

	//Initialize gobal Count of  in each class
	global.InitGlobal()

	//initailize gin and its router path
	GinApp = gin.Default()
	api.ApplyRoutes(GinApp)

}
