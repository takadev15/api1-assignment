package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/go-restapi/internal/database"
	"github.com/takadev15/go-restapi/internal/handlers"
)

func OrderRoutes() *gin.Engine {
	routes := gin.Default()
	db := database.GetDB()

	orderHandler := &handlers.OrderHandler{Connect: db}
  orderRoutes := routes.Group("/order")
    {
	    orderRoutes.GET("/", orderHandler.GetAllOrder)
      orderRoutes.GET("/:id", orderHandler.GetOrder)
      orderRoutes.POST("/", orderHandler.CreateOrder)
      orderRoutes.PUT("/", orderHandler.UpdateOrder)
      orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
    }

  return routes
}
