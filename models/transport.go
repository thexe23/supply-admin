package models

import (
	"gorm.io/gorm"
	"time"
)

type Transport struct {
	ID       int64     `json:"id"`
	SourceID int64     `json:"sourceId"`
	TargetID int64     `json:"targetId"`
	Item     string    `json:"item"`
	Quantity int       `json:"quantity"`
	CreateAt time.Time `json:"createAt"`
	Status   int       `json:"status"`
}

func AddTransport(sourceID, targetID int64, item string, quantity int) (id int64, err error) {
	status := 10
	if sourceID == 9999 {
		status = 20
	}
	var transport = Transport{
		SourceID: sourceID,
		TargetID: targetID,
		Item:     item,
		Quantity: quantity,
		Status: status,
	}
	tx := db.Begin()
	tx.Table("transport").Select("SourceID", "TargetID", "Item", "Quantity", "Status").Create(&transport)
	if sourceID != 9999 {
		tx.Table("item").Where("name = ? AND market_id = ?", transport.Item, transport.SourceID).Update("stock", gorm.Expr("stock - ?", transport.Quantity))
	}
	err = tx.Commit().Error
	if err != nil || transport.ID == 0 {
		tx.Rollback()
		return 0, err
	}
	return transport.ID, err
}

func UpdateTransportStatus(id int64, status int) error {
	tx := db.Begin()
	tx.Table("transport").Where("id = ?", id).Update("status", status)
	if status == 30 {
		var ts Transport
		tx.Table("transport").Select("target_id", "item", "quantity").Where("id =?", id).Find(&ts)
		tx.Table("item").Where("market_id = ? AND name = ?", ts.TargetID, ts.Item).Update("stock", gorm.Expr("stock + ?", ts.Quantity))
	}
	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func GetTransportForSource(userID int64) (error, []Transport) {
	var transports []Transport
	err := db.Table("transport").Where("source_id = ?", userID).Find(&transports).Error
	if err != nil {
		return err, nil
	}
	return nil, transports
}


func GetTransportForTarget(userID int64) (error, []Transport) {
	var transports []Transport
	err := db.Table("transport").Where("target_id = ?", userID).Find(&transports).Error
	if err != nil {
		return err, nil
	}
	return nil, transports
}


func GetTransport() (error, []Transport) {
	var transports []Transport
	err := db.Table("transport").Find(&transports).Error
	if err != nil {
		return err, nil
	}
	return nil, transports
}

