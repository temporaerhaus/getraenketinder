package api

import (
	"github.com/gin-gonic/gin"
)

// SetupEndpoints initializes a gin Engine and populates it with endpoint handler functions
func SetupEndpoints(router *gin.Engine) {
	v0 := router.Group("/api")
	{
		getr := v0.Group("/getraenk")
		{
			getr.PUT("", putGetraenk)
			getr.GET("", getGetraenk)
			getr.GET("/:uuid", getrGetraenkByUUID)
			getr.PATCH(":/uuid", patchGetraenkByUUID)
			getr.DELETE("/:uuid", deleteGetraenkByUUID)
			getr.POST("/:uuid/dislike", postDislikeByUUID)
			getr.POST("/:uuid/like", postLikeByUUID)
			getr.POST("/:uuid/superlike", postSuperlikeByUUID)
		}
	}

}

func postSuperlikeByUUID(context *gin.Context) {

}

func postLikeByUUID(context *gin.Context) {

}

func postDislikeByUUID(context *gin.Context) {

}

func deleteGetraenkByUUID(context *gin.Context) {

}

func patchGetraenkByUUID(context *gin.Context) {

}

func getrGetraenkByUUID(context *gin.Context) {

}

func getGetraenk(context *gin.Context) {

}

func putGetraenk(context *gin.Context) {

}
