package app

import (
	"model"
	"net/http"
	"resources"
	"time"
	"github.com/gin-gonic/gin"

)

type OrderController struct {
	service OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		service: newOrderService(),
	}
}

type OrderOut struct {
	ID           uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []ItemOut `gorm:"foreignKey:OrderID"`
}

type ItemOut struct {
	ItemID      uint   `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

func (controller *OrderController) addOrder(context *gin.Context)  {
	service := controller.service
	var request resources.InputOrder
	err := context.ShouldBind(&request)
	if context.Request.Method == "PUT" && request.Id == 0 {
		response := resources.JsonResponse("bad request", http.StatusBadRequest, "error", "ID is Mandatory")
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := resources.JsonResponse("bad request", http.StatusBadRequest, "error", "Something When Wrong.")
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var order model.Order
	if request.Id != 0 {
		order, _ = service.getOrderDetailById(request.Id, false)
		if order.Id == 0 {
			response := resources.JsonResponse("Order not found",http.StatusBadRequest, "error", "")
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		err := service.addOrder(&order, request)
		if err != nil {
			response := resources.JsonResponse("Update Order Failed", http.StatusInternalServerError, "erroe", "")
			context.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
	} else {
		order = model.Order{}
		err = service.addOrder(&order, request)
		if err != nil {
			response := resources.JsonResponse("Update Order Failed", http.StatusInternalServerError, "error", "")
			context.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
	}
	orderResult, err := service.getOrderDetailById(uint(order.Id), true)
	filteredOrder := OrderOut{}
	filteredOrder.ID = orderResult.ID
	filteredOrder.CustomerName = orderResult.CustomerName
	filteredOrder.OrderedAt = orderResult.OrderedAt
	filteredOrder.Items = []ItemOut{}
	for _, itemValue := range orderResult.Items {
		eachItem := ItemOut{}
		eachItem.ItemID = itemValue.ID
		eachItem.ItemCode = itemValue.ItemCode
		eachItem.Description = itemValue.Description
		eachItem.Quantity = itemValue.Quantity
		eachItem.OrderID = itemValue.OrderID

		filteredOrder.Items = append(filteredOrder.Items, eachItem)
	}

	response :=resources.JsonResponsePagination("Create Order Success", http.StatusOK, 0, 0, 0, filteredOrder)
	context.JSON(http.StatusOK, response)
}

