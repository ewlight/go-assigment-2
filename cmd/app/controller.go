package app

import (
	"context"
	"model"
	"net/http"
	"resources"
	"strconv"
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

	response := resources.JsonResponse("Create Order Success", http.StatusOK, "Success", filteredOrder)
	context.JSON(http.StatusOK, response)
}

func (controller *OrderController) getOrderList(context *gin.Context) {
	service := controller.service
	result, err, count := service.getOrderList()
	if err != nil {
		response := resources.JsonResponse("Something When Wrong", http.StatusInternalServerError, "error", "")
		context.AbortWithStatusJSON(http.StatusInternalServerError, response)
	}

	orderList := []OrderOut{}
	for _, value := range result {
		eachOrder := OrderOut{}
		eachOrder.ID = value.ID
		eachOrder.CustomerName = value.CustomerName
		eachOrder.OrderedAt = value.OrderedAt
		eachOrder.Items = []ItemOut{}
		for _, itemValue := range value.Items {
			eachItem := ItemOut{}
			eachItem.ItemID = itemValue.ID
			eachItem.ItemCode = itemValue.ItemCode
			eachItem.Description = itemValue.Description
			eachItem.Quantity = itemValue.Quantity
			eachItem.OrderID = itemValue.OrderID

			eachOrder.Items = append(eachOrder.Items, eachItem)
		}
		orderList = append(orderList, eachOrder)
	}

	response := resources.JsonResponsePagination("Success Retrieve Order", http.StatusOK, 0,0,int(count), orderList)
	context.JSON(http.StatusOK, response)
}

func (controller *OrderController) getOrderDetail(context *gin.Context) {
	service := controller.service
	orderId, _ := strconv.Atoi(context.Param("order_id"))
	orderResult, err := service.getOrderDetailById(uint(orderId), true)
	if err != nil {
		response := resources.JsonResponse("Order Not Found", http.StatusNotFound, "error", "")
		context.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}
	eachOrder := OrderOut{}
	eachOrder.ID = orderResult.ID
	eachOrder.CustomerName = orderResult.CustomerName
	eachOrder.OrderedAt = orderResult.OrderedAt
	eachOrder.Items = []ItemOut{}
	for _, itemValue := range orderResult.Items {
		eachItem := ItemOut{}
		eachItem.ItemID = itemValue.ID
		eachItem.ItemCode = itemValue.ItemCode
		eachItem.Description = itemValue.Description
		eachItem.Quantity = itemValue.Quantity
		eachItem.OrderID = itemValue.OrderID

		eachOrder.Items = append(eachOrder.Items, eachItem)
	}
	// \ Converting result to json

	response := resources.JsonResponse("Success Retrieve Order", http.StatusOK, "Success", eachOrder)
	context.JSON(http.StatusOK, response)
}

func (controller *OrderController) deleteOrder(context *gin.Context) {
	id, _ := strconv.Atoi(c.Param("order_id"))
	if id == 0 {
		response := resources.JsonResponse("order_id is required", http.StatusBadRequest, "error", "")
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	service := controller.service
	delete := service.deleteOrder(id)
	if delete != nil {
		response := resources.JsonResponse("Error on Deleting Order", http.StatusInternalServerError, "error", "")
		context.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := resources.JsonResponse("Deleting Order Success", http.StatusOK,"Success","")
	context.JSON(http.StatusOK, response)
}



