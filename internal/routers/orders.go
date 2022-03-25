package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/go-restapi/internal/database"
	"github.com/takadev15/go-restapi/internal/handlers"
  docs "github.com/takadev15/go-restapi/docs"
  swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)

func OrderRoutes() *gin.Engine {
	routes := gin.Default()
	db := database.GetDB()
  docs.SwaggerInfo.BasePath = "/order"
	orderHandler := &handlers.OrderHandler{Connect: db}
  orderRoutes := routes.Group("/order")
    {
	    orderRoutes.GET("/", orderHandler.GetAllOrder)
      orderRoutes.GET("/:id", orderHandler.GetOrder)
      orderRoutes.POST("/", orderHandler.CreateOrder)
      orderRoutes.PUT("/", orderHandler.UpdateOrder)
      orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
    }
  routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

  return routes
}
