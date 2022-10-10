package app

import "github.com/gin-gonic/gin"

func initRouters() {
	orderController := NewOrderController()
	server := gin.Default()
	server.POST("/order", orderController.addOrder)
	server.PUT("/order", orderController.addOrder)
	server.GET("/order", orderController.getOrderList)
	server.GET("/order/:order_id", orderController.getOrderDetail)
	server.DELETE("/order/:order_id", orderController.deleteOrder)
	server.Run(":8080")
}
