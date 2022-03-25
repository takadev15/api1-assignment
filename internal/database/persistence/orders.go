package persistence

import (
	"github.com/takadev15/go-restapi/internal/models"
	"gorm.io/gorm"
)


func GetAllOrder(db *gorm.DB) ([]models.Order, error) {
  var orders []models.Order
  res := db.Find(&orders)
  if res.Error != nil {
    return nil, res.Error
  } else {
      if res.RowsAffected <= 0 {
        return nil, res.Error
      } else {
        return orders, res.Error
      }
  }
}

func GetOrder(id int, db *gorm.DB) (models.Order, error){
  var order models.Order
  err := db.Where("id = ?", id).First(&order).Error
  if err != nil {
    return models.Order{}, err
  }
  return order, err
}

func CreateOrder(data *models.Order, db *gorm.DB) error {
  res := db.Create(&data)
  if res.Error != nil {
    return res.Error
  }
  return nil
}

func UpdateOrder(id int, data *models.Order, db *gorm.DB) (models.Order, error){
  var order models.Order

  err := db.Preload("Items").First(&order, id).Error
  if err != nil {
    return models.Order{}, err
  }
  err = db.Model(&order).Updates(&data).Error
  if err != nil {
    return models.Order{}, err
  }
  return order, err
}

func DeleteOrder(id int, db *gorm.DB) (error){
  var order models.Order
  err := db.First(&order, id).Error
  if err != nil {
    return err
  }
  err = db.Delete(&order).Error
  if err != nil {
    return err
  } else {
    return nil
  }
}

