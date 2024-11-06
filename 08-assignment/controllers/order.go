package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/retry19/challenge-hacktiv8/08-assignment/database"
	"github.com/retry19/challenge-hacktiv8/08-assignment/models"
	"gorm.io/gorm"
)

// @GetOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags Orders
// @Produce json
// @Success 200 {object} models.ResponseData{data=[]models.Order}
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	var orders []models.Order

	err := database.PgClient.Preload("Items").Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &models.ResponseData{
		Status: true,
		Data:   orders,
	})
}

// @CreateOrder godoc
// @Summary Create an order
// @Description Create an order
// @Tags Orders
// @Produce json
// @Param bodyPayload body models.CreateOrderDto true "Create order dto"
// @Success 200 {object} models.ResponseData{data=models.Order}
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var bodyPayload models.CreateOrderDto

	if err := c.ShouldBindJSON(&bodyPayload); err != nil {
		c.JSON(http.StatusBadRequest, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	order := models.Order{
		CustomerName: bodyPayload.CustomerName,
		OrderedAt:    bodyPayload.OrderedAt,
	}

	for _, item := range bodyPayload.Items {
		orderItem := models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}
		order.Items = append(order.Items, orderItem)
	}

	if err := database.PgClient.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &models.ResponseData{
		Status: true,
		Data:   order,
	})
}

// @UpdateOrder godoc
// @Summary Update an order
// @Description Update an order by id
// @Tags Orders
// @Produce json
// @Param id path int true "Order id"
// @Param bodyPayload body models.UpdateOrderDto true "Update order dto"
// @Success 200 {object} models.ResponseData{data=models.Order}
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /orders/{id} [put]
func UpdateOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	var order models.Order
	err = database.PgClient.Preload("Items").First(&order, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	var bodyPayload models.UpdateOrderDto

	if err := c.ShouldBindJSON(&bodyPayload); err != nil {
		c.JSON(http.StatusBadRequest, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	err = database.PgClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order).Updates(&models.Order{
			CustomerName: bodyPayload.CustomerName,
		}).Error; err != nil {
			return err
		}

		items := []models.Item{}

		fmt.Println("items", items)

		for _, item := range bodyPayload.Items {
			items = append(items, models.Item{
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
				OrderId:     order.Id,
			})
		}

		if err := tx.Delete(&order.Items).Error; err != nil {
			return err
		}

		if err := tx.CreateInBatches(items, 100).Error; err != nil {
			return err
		}

		order.Items = items

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &models.ResponseData{
		Status: true,
		Data:   order,
	})
}

// @DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order by id
// @Tags Orders
// @Param id path int true "Order id"
// @Success 200 {object} models.ResponseData
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /orders/{id} [delete]
func DeleteOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	order := models.Order{}
	err = database.PgClient.Preload("Items").First(&order, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	err = database.PgClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&order.Items).Error; err != nil {
			return err
		}

		if err := tx.Delete(&order).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.ResponseError{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &models.ResponseData{
		Status: true,
	})
}
