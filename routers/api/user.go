package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"supply-admin/service"
	errMsg "supply-admin/service/error"
	"supply-admin/service/user_service"
)

func GetUsers(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	err, users := user_service.GetUsers()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_USER_FAIL, err)
		return
	}

	appG.Response(http.StatusOK, errMsg.SUCCESS, users)
}

func GetUserByID(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	id, ok := c.Params.Get("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil || !ok {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	err, user := user_service.GetUserByID(int64(userID))
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_USER_FAIL, err)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, user)
}

type updateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	ImgUrl   string `json:"imgUrl"`
}

func UpdateUser(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	var user = &updateUserReq{}
	err := c.ShouldBindJSON(&user)
	id, ok := c.Params.Get("user_id")
	userID, _ := strconv.Atoi(id)
	if err != nil || !ok {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	userService := user_service.User{
		ID:       int64(userID),
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		ImgUrl:   user.ImgUrl,
	}
	err = userService.Update()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_REGISTER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errMsg.SUCCESS, nil)
}

func DeleteUser(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	id, ok := c.Params.Get("user_id")
	userID, err := strconv.Atoi(id)
	if err != nil && !ok {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	userService := user_service.User{
		ID:       int64(userID),
	}
	err = userService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, errMsg.SUCCESS, nil)
}
