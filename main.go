package main

import (
	"fmt"

	"apiv0.1/controller"
	"apiv0.1/storage"
	"github.com/labstack/echo/v4"
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
	e.GET("/api/v0/product", controller.Index)
	e.GET("/api/v0/product/:id", controller.Show)
	e.POST("/api/v0/product", controller.Store)
	e.PUT("/api/v0/product/:id", controller.Update)
	e.DELETE("/api/v0/product/:id", controller.Destroy)

	fmt.Println("Servidor started!!")
	e.Logger.Fatal(e.Start(":999"))
}
