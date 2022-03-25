package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/go-restapi/internal/database/persistence"
	"github.com/takadev15/go-restapi/internal/models"
	"gorm.io/gorm"
)

type OrderHandler struct {
  Connect *gorm.DB
}

type ResponseItems struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ResponseOrder struct {
	OrderedAt    time.Time     `json:"ordered_at"`
	CustomerName string        `json:"customer_name"`
	Items        ResponseItems `json:"items"`
}
// @BasePath /order

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description get all order
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} GetAllOrder()
// @Router /order [get]
func(db OrderHandler) GetAllOrder(c *gin.Context) {
  var (
    result gin.H
    )
  orderResult:= make([]ResponseOrder, 10)
  res, err := persistence.GetAllOrder(db.Connect)
  for i := range res {
    orderResult[i].OrderedAt = res[i].CreatedAt
    orderResult[i].CustomerName = res[i].CustomerName
    orderResult[i].Items.ItemCode = res[i].Item.ItemsCode
    orderResult[i].Items.Description = res[i].Item.Description
    orderResult[i].Items.Quantity = res[i].Item.Quantity
  }
  if err != nil {
    result = gin.H{
      "message": err,
    }
  }
  result = gin.H{
    "result": orderResult,
  }
  c.JSON(http.StatusOK, result)
}

func (db OrderHandler) GetOrder(c *gin.Context) {
  var (
    result gin.H
    orderResult ResponseOrder
    )
  inputId := c.Param("id")
  orderId, _ := strconv.Atoi(inputId)
  res, err := persistence.GetOrder(orderId, db.Connect)
  {
    orderResult.OrderedAt = res.CreatedAt
    orderResult.CustomerName = res.CustomerName
    orderResult.Items.ItemCode = res.Item.ItemsCode
    orderResult.Items.Description = res.Item.Description
    orderResult.Items.Quantity = res.Item.Quantity
  }
  if err != nil {
    result = gin.H{
      "message": err,
    }
  }
  result = gin.H{
    "result": orderResult,
  }
  c.JSON(http.StatusOK, result)
}

func (db OrderHandler) CreateOrder(c *gin.Context) {
  var (
    order models.Order
    result gin.H
    )
  if err := c.ShouldBindJSON(&order); err != nil {
    c.AbortWithStatus(http.StatusBadRequest)
  }
  err := persistence.CreateOrder(&order, db.Connect)
  if err != nil {
    result = gin.H{
      "message": err,
    }
  }
  result = gin.H{
    "result": order,
  }
  c.JSON(http.StatusOK, result)
}

func (db OrderHandler) UpdateOrder(c *gin.Context) {
  var (
    order models.Order
    result gin.H
    )
  if err := c.ShouldBindJSON(&order); err != nil {
    c.AbortWithStatus(http.StatusBadRequest)
  }
  orderId := int(order.ID)
  _, err := persistence.UpdateOrder(orderId, &order, db.Connect)
  if err != nil {
    result = gin.H{
      "message": err,
    }
  }
  result = gin.H{
    "result": order,
  }
  c.JSON(http.StatusOK, result)
}

func (db OrderHandler) DeleteOrder(c *gin.Context) {
  var result gin.H
  inputId := c.Param("id")
  id, _ := strconv.Atoi(inputId)
  err := persistence.DeleteOrder(id, db.Connect)
  if err != nil {
    result = gin.H{
      "message" : err,
    }
  }
  result = gin.H{
    "message": "berhasil dihapus",
  }
  c.JSON(http.StatusOK, result)
}

