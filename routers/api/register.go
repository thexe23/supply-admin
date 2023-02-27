package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"supply-admin/service"
	errMsg "supply-admin/service/error"
	"supply-admin/service/user_service"
)

type register struct {
	Username string `validate:"gt=0" ,json:"username"`
	Password string `validate:"gt=0" ,json:"password"`
	Phone    string `validate:"gt=0" ,json:"phone"`
	Role     string `validate:"gt=0" ,json:"role"`
	OrgID   string `validate:"gt=0" ,json:"orgId"`
}

func Register(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	var register = &register{}
	err := c.ShouldBindJSON(&register)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	var validate = validator.New()
	err = validate.Struct(register)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	role, _ := strconv.Atoi(register.Role)
	orgId, _ := strconv.Atoi(register.OrgID)
	marketID, _ := user_service.MarketMap[int64(orgId)]
	userService := user_service.User{
		Username: register.Username,
		Password: register.Password,
		Phone:    register.Phone,
		Role:     role,
		OrgID:    int64(orgId),
		MarketID: marketID,
	}
	id, err := userService.Add()
	if err != nil || id == 0 {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_REGISTER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errMsg.SUCCESS, map[string]int64{
		"id": id,
	})
}
