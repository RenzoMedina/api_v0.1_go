package main

import (
	"fmt"

	"apiv0.1/controller"
	"apiv0.1/storage"
	"github.com/labstack/echo/v4"

	"github.com/rs/cors"
)

func main() {
	/*
		? Conection of the database
	*/
	storage.NewConnection()

	e := echo.New()
	/*
	* Routes
	 */
	e.GET("v0/product", controller.Index)
	e.GET("v0/product/:id", controller.Show)
	e.POST("v0/product", controller.Store)
	e.PUT("v0/product/:id", controller.Update)
	e.DELETE("v0/product/:id", controller.Destroy)

	/*
		? Cors
	*/
	c := cors.Default().Handler(e)

	fmt.Println("Servidor started!!")
	e.Logger.Fatal(e.Start(":999"), c)
}
