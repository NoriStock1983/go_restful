package commons

import (
	"GOTEST/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	mastermaintenancerouter := router.Group("mastermaintenance")
	{
		usermst := mastermaintenancerouter.Group("usermst")
		{
			usermst.GET("/allusermst", controllers.Allusermst)
			usermst.GET("/getbyusercd/:usercd", controllers.GetByUsercd)
		}
	}

	return router
}
