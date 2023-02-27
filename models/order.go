package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	MarketID    int64     `json:"marketId"`
	OrgID       int       `json:"orgId"`
	Items       string    `json:"items"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
	OrderStatus int       `json:"orderStatus"`
}

type CartItem struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	ImgUrl string `json:"imgUrl"`
	Number int    `json:"number"`
}

func AddOrder(userID, marketID int64, orgId int, items string, amount int) (id int64, err error) {
	var order = Order{
		UserID:   userID,
		MarketID: marketID,
		OrgID:    orgId,
		Items:    items,
		Amount:   amount,
	}
	tx := db.Begin()
	tx.Table("order").Select("UserID", "MarketID", "OrgID", "Items", "Amount").Create(&order)
	var cartItem []CartItem
	json.Unmarshal([]byte(items), &cartItem)
	for _, v := range cartItem {
		tx.Table("item").Where("id = ?", v.ID).Update("stock", gorm.Expr("stock - ?", v.Number))
	}
	err = tx.Commit().Error
	if err != nil || order.ID == 0 {
		tx.Rollback()
		return 0, err
	}
	return order.ID, err
}

func UpdateStatus(id int64, status int) error {
	err := db.Table("order").Where("id = ?", id).Update("order_status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrdersForUser(userID int64) (error, []Order) {
	var orders []Order
	err := db.Table("order").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return err, nil
	}
	return nil, orders
}

func GetOrdersForMarket(userID int64) (error, []Order) {
	var orders []Order
	err := db.Table("order").Where("market_id = ?", userID).Find(&orders).Error
	if err != nil {
		return err, nil
	}
	return nil, orders
}

func GetOrdersByOrg(orgID int) (error, []Order) {
	var orders []Order
	err := db.Table("order").Where("org_id = ?", orgID).Find(&orders).Error
	if err != nil {
		return err, nil
	}
	return nil, orders
}
