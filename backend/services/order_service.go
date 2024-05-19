package services

import (
	"errors"
)

type OrderService interface {
	GetOrderStatus(orderID int) (string, error)
	InitiateReturn(orderID int) (string, error)
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}

func (s *orderService) GetOrderStatus(orderID int) (string, error) {
	// Simulate fetching order status
	if orderID <= 0 {
		return "", errors.New("invalid order ID")
	}
	return "Your order is being processed.", nil
}

func (s *orderService) InitiateReturn(orderID int) (string, error) {
	// Simulate initiating a return
	if orderID <= 0 {
		return "", errors.New("invalid order ID")
	}
	return "Your return request has been initiated.", nil
}
