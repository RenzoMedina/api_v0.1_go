package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"apiv0.1/model"
	"apiv0.1/storage"
	"github.com/labstack/echo/v4"
)

func Store(c echo.Context) error {
	product := &model.Product{
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	}
	mysqlser := storage.NewMySQL(storage.Pool())
	newmysql := model.NewServices(mysqlser)

	if err := newmysql.Create(product); err != nil {
		log.Fatalf("model.Create() %v", err)
	}
	return c.JSON(http.StatusOK, product)
}

func Update(c echo.Context) error {

	//convert type string to int
	id := c.Param("id")
	va, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	product := &model.Product{
		ID:        uint(va),
		Title:     c.FormValue("title"),
		Body:      c.FormValue("body"),
		Update_At: time.Now(),
	}
	mysqlser := storage.NewMySQL(storage.Pool())
	newmysql := model.NewServices(mysqlser)

	if err := newmysql.Update(product); err != nil {
		log.Fatalf("model.Update() %v", err)
	}
	return c.JSON(http.StatusOK, product)
}

func Destroy(c echo.Context) error {

	//convert type string to int
	id := c.Param("id")
	va, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	mysqlser := storage.NewMySQL(storage.Pool())
	newmysql := model.NewServices(mysqlser)

	if err := newmysql.Delete(uint(va)); err != nil {
		log.Fatalf("model.Delete() %v", err)
	}
	return c.JSON(http.StatusOK, "Data delete ok!")

}
