package auth_service

import "supply-admin/models"

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Auth) Check() (bool, int64,  error) {
	return models.CheckAuth(a.Username, a.Password)
}
