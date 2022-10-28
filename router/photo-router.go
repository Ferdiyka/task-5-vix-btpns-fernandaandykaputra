package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/controllers"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/helpers"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/middlewares"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/models"
)

var (
	photoRepo       models.PhotoRepo            = models.NewPhotoRepo(db)
	photoHelper     helpers.PhotoHelper         = helpers.NewPhotoHelper(photoRepo)
	jwtHelper       helpers.JWTHelper           = helpers.NewJWTHelper()
	photoController controllers.PhotoController = controllers.NewPhotoController(photoHelper, jwtHelper)
)

func PhotoRouter() {
	router := gin.Default()

	photoRoutes := router.Group("api/photo", middlewares.AuthJWT(jwtHelper))
	{
		photoRoutes.POST("/photos", photoController.Insert)
		photoRoutes.PUT("/:photoId", photoController.Update)
		photoRoutes.GET("/photos", photoController.GetByID)
		photoRoutes.DELETE("/:photoId", photoController.Delete)
	}
	router.Run()
}