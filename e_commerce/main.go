package main

import (
	"log"
	"os"

	"lib-ser/controllers"
	"lib-ser/database"
	"lib-ser/middleware"
	"lib-ser/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	gin_router := gin.New()

	routes.UserRoutes(gin_router)

	gin_router.Use(gin.Logger())
	gin_router.Use(middleware.Authentication())

	gin_router.GET("/addtocart", app.AddToCart())
	gin_router.GET("/removeitem", app.RemoveItem())
	gin_router.GET("/listcart", controllers.GetItemFromCart())
	gin_router.POST("/addaddress", controllers.AddAddress())
	gin_router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	//gin_router.PUT("/editworkaddress", controllers.EditWorkAddress())
	gin_router.GET("/deleteaddresses", controllers.DeleteAddress())
	gin_router.GET("/cartcheckout", app.BuyFromCart())
	gin_router.GET("/instantbuy", app.InstantBuy())
	log.Fatal(gin_router.Run(":" + port))
}
