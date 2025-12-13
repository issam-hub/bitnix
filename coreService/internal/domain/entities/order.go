package entities

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderPaid      OrderStatus = "paid"
	OrderFailed    OrderStatus = "failed"
	OrderCancelled OrderStatus = "cancelled"
	OrderRefunded  OrderStatus = "refunded"
)

type Order struct {
	OrderID     uuid.UUID
	UserID      uuid.UUID
	Items       []Game
	TotalAmount money.Money
	Status      OrderStatus
}
