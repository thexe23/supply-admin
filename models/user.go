package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`
	OrgID    int64  `json:"orgId"`
	MarketID int64  `json:"marketId"`
	ImgUrl   string `json:"imgUrl"`
}

func AddUser(username, password, phone string, role int, orgID, marketID int64) (id int64, err error) {
	var user = User{
		Username: username,
		Password: password,
		Phone:    phone,
		Role:     role,
		OrgID:    orgID,
		MarketID: marketID,
	}
	err = db.Table("user").Select("Username", "Password", "Phone", "Role", "OrgID", "MarketID", "ImgUrl").Create(&user).Error
	if err != nil || user.ID == 0 {
		return 0, err
	}
	return user.ID, nil
}

func GetUsers() (error, []User) {
	var users []User
	err := db.Table("user").Find(&users).Error
	if err != nil {
		return err, nil
	}
	return nil, users
}

func UpdateUser(id int64, username, password, phone, imgUrl string) error {
	return db.Table("user").Where("id = ?", id).Updates(User{
		Username: username,
		Password: password,
		Phone:    phone,
		ImgUrl:   imgUrl,
	}).Error
}

func DeleteUser(id int64) error {
	return db.Table("user").Delete(&User{}, id).Error
}

func GetUserByID(id int64) (error, *User) {
	var user User
	err := db.Table("user").Where("id = ?", id).First(&user).Error
	if err != nil {
		return err, nil
	}
	return nil, &user
}
