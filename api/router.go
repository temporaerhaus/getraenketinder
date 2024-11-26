package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func putGetraenk(context *gin.Context, db *sql.DB) {
	//Extracting context body into object and generating UUID
	var uploadGetraenk UploadGetraenk
	if err := context.ShouldBind(&uploadGetraenk); err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}
	if uploadGetraenk.UUID == ""{
		uploadGetraenk.UUID = uuid.New().String()
	}
	//TODO: Bilder in irgendwo speichern und sich die URLs merken, dann das UploadGetraenk in ein Getraenk überführen
	
	var getraenk Getraenk
	
	//Writing the getraenk in the DB
	if err := CreateNewGetraenk(db, &getraenk); err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
	}
}