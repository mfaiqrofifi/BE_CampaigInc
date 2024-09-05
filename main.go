package main

import (
	config "Campign/database"
	"Campign/domain/item/controller"
	"Campign/domain/item/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	route := echo.New()

	itemControllerUser := controller.NewControllerUser(db)
	apiV1 := route.Group("api/v1/")
	apiV1.POST("user/register", itemControllerUser.Register)
	apiV1.POST("user/login", itemControllerUser.Login)
	apiV1.GET("user/get", itemControllerUser.GetAllUse)

	protected := route.Group("api/v1/")
	protected.Use(middleware.VeryfyJWT)

	itemControllerCampaign := controller.NewControllerCampaign(db)
	protected.POST("campaign/create", itemControllerCampaign.Create)
	protected.PUT("campaign/update/:id", itemControllerCampaign.Update)
	protected.DELETE("campaign/delete/:id", itemControllerCampaign.Delete)
	protected.GET("campaign/get", itemControllerCampaign.GetAll)
	protected.GET("campaign/detail", itemControllerCampaign.GetById)

	route.Start(":9000")
}
