package order_service

import (
	"supply-admin/models"
)

type Order struct {
	UserID   int64  `json:"user_id"`
	MarketID int64  `json:"market_id"`
	OrgID    int  `json:"org_id"`
	Items    string `json:"items"`
	Amount   int    `json:"amount"`
}

var StatusMap = map[int]string{
	-1: "已取消",
	10: "待支付",
	20: "待发货",
	30: "配送中",
	40: "待取货",
	50: "已完成",
}

func (o *Order) Add() (int64, error) {
	return models.AddOrder(o.UserID, o.MarketID, o.OrgID, o.Items, o.Amount)
}

func GetForUser(userId int64) (error, []models.Order) {
	return models.GetOrdersForUser(userId)
}

func GetForMarket(marketID int64) (error, []models.Order) {
	return models.GetOrdersForMarket(marketID)
}

func GetByOrg(orgID int) (error, []models.Order) {
	return models.GetOrdersByOrg(orgID)
}

type OrderStatus struct {
	ID          int64 `json:"id"`
	OrderStatus int   `json:"order_status"`
}

func (o *OrderStatus) UpdateStatus() error {
	return models.UpdateStatus(o.ID, o.OrderStatus)
}
