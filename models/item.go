package models

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

func AddItem(marketID int64, name string, price, stock, category int, imgUrl string) (id int64, err error) {
	var item = Item{
		MarketID: marketID,
		Name:     name,
		Price:    price,
		Stock:    stock,
		Category: category,
		ImgUrl:   imgUrl,
	}
	err = db.Table("item").Select("MarketID", "Name", "Price", "Stock", "Category", "ImgUrl").Create(&item).Error
	if err != nil || item.ID == 0 {
		return 0, err
	}
	return item.ID, nil
}

func UpdateItem(id int64, name string, price, stock, category int, imgUrl string) error {
	err := db.Table("item").Where("id = ?", id).Updates(Item{Name: name, Price: price, Stock: stock, Category: category, ImgUrl: imgUrl}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetItemByID(marketID int64) (error, []Item) {
	var res []Item
	err := db.Table("item").Where("market_id = ?", marketID).Order("on_sale desc").Find(&res).Error
	if err != nil {
		return err, nil
	}
	return nil, res
}

func GetAllItem() (error, []Item) {
	var res []Item
	err := db.Table("item").Find(&res).Error
	if err != nil {
		return err, nil
	}
	return nil, res
}

func GetItemOnSale(marketID int64) (error, []Item) {
	var res []Item
	err := db.Table("item").Where("market_id = ? AND on_sale = 1", marketID).Find(&res).Error
	if err != nil {
		return err, nil
	}
	return nil, res
}

func ToggleItem(id int64, onSale bool) error {
	return db.Table("item").Where("id = ?", id).Update("on_sale", onSale).Error
}


