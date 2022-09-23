package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanramli/hacktiv8-assignment-2/httpserver/controllers"
)

type router struct {
	router *gin.Engine

	order *controllers.OrderController
}

func NewRouter(r *gin.Engine, order *controllers.OrderController) *router {
	return &router{
		router: r,
		order:  order,
	}
}

func (r *router) Start(port string) {
	r.router.GET("/orders", r.order.GetAllOrder)
	r.router.DELETE("/orders/:id", r.order.DeleteOrder)
	r.router.POST("/orders", r.order.CreateOrder)
	r.router.PUT("/orders", r.order.UpdateOrder)
	r.router.Run(port)
}
