package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"supply-admin/service"
	"supply-admin/service/auth_service"
	errMsg "supply-admin/service/error"
	"supply-admin/util"
)

type auth struct {
	Username string `validate:"gt=0" ,json:"username"`
	Password string `validate:"gt=0" ,json:"password"`
}

func GetAuth(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	var auth = &auth{}
	c.BindJSON(&auth)
	var validate = validator.New()
	err := validate.Struct(auth)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	authService := auth_service.Auth{
		Username: auth.Username,
		Password: auth.Password,
	}
	isExist, userID, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_AUTH, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, errMsg.ERROR_AUTH, nil)
		return
	}
	token, err := util.CreateToken(userID)
	if err != nil {
		appG.Response(http.StatusUnauthorized, errMsg.ERROR_AUTH_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, map[string]string {
		"token": token,
	})
}
