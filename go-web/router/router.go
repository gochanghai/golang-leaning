package router

import (
	"github.com/gin-gonic/gin"
	. "golang-leaning/go-web/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getUsers", Users)

	router.POST("/addUser", Store)

	router.PUT("/updateUser/:id", Update)

	router.DELETE("/deleteUser/:id", Destroy)

	return router
}
