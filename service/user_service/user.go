package user_service

import "supply-admin/models"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`
	OrgID    int64  `json:"org_id"`
	MarketID int64  `json:"market_id"`
	ImgUrl   string  `json:"imgUrl"`
}

var MarketMap = map[int64]int64{
	1001: 10008,
	1002: 10008,
	1003: 10009,
	1004: 10009,
}

func (u *User) Add() (int64, error) {
	return models.AddUser(u.Username, u.Password, u.Phone, u.Role, u.OrgID, u.MarketID)
}

func (u *User) Update() error {
	return models.UpdateUser(u.ID, u.Username, u.Password, u.Phone, u.ImgUrl)
}

func (u *User) Delete() error {
	return models.DeleteUser(u.ID)
}

func GetUsers() (error, []models.User) {
	return models.GetUsers()
}

func GetUserByID(id int64) (error, *models.User) {
	return models.GetUserByID(id)
}
