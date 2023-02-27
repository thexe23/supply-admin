package item_service

import "supply-admin/models"

type Item struct {
	ID       int64  `json:"id"`
	MarketID int64  `json:"marketId"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Category int    `json:"category"`
	ImgUrl   string `json:"imgUrl"`
	OnSale   bool   `json:"onSale"`
}

func (i *Item) Add() (int64, error) {
	return models.AddItem(i.MarketID, i.Name, i.Price, i.Stock, i.Category, i.ImgUrl)
}

func (i *Item) Update() error {
	return models.UpdateItem(i.ID, i.Name, i.Price, i.Stock, i.Category, i.ImgUrl)
}

func (i *Item) Toggle() error {
	return models.ToggleItem(i.ID, i.OnSale)
}

func GetByID(marketID int64) (error, []models.Item) {
	return models.GetItemByID(marketID)
}

func GetOnSale(marketID int64) (error, []models.Item) {
	return models.GetItemOnSale(marketID)
}

func GetAll() (error, []models.Item) {
	return models.GetAllItem()
}