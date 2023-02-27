package models

import "gorm.io/gorm"

type Auth struct {
	ID int64 `gorm:"primary_key", json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, int64, error){
	var auth Auth
	err := db.Table("user").Select("id", "role").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, 0, err
	}

	if auth.ID > 0 {
		return true, auth.ID, nil
	}

	return false, 0, nil
}
