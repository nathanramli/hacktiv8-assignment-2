package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanramli/hacktiv8-assignment-2/config"
	"github.com/nathanramli/hacktiv8-assignment-2/httpserver"
	"github.com/nathanramli/hacktiv8-assignment-2/httpserver/controllers"
	"github.com/nathanramli/hacktiv8-assignment-2/httpserver/repositories/gorm"
	"github.com/nathanramli/hacktiv8-assignment-2/httpserver/services"
)

func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	itemRepo := gorm.NewItemRepo(db)
	orderRepo := gorm.NewOrderRepo(db)
	orderSvc := services.NewOrderSvc(orderRepo, itemRepo)
	orderHandler := controllers.NewOrderController(orderSvc)

	app := httpserver.NewRouter(router, orderHandler)
	app.Start(":8000")
}
