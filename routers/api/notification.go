package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"supply-admin/service"
	errMsg "supply-admin/service/error"
	"supply-admin/service/notification_service"
)

func AddNotification(c *gin.Context) {
	appG := service.Gin{
		Ctx: c,
	}
	msg := notification_service.Notification{}
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, err)
		return
	}
	err = notification_service.Add(msg)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR, err)
		return
	}

	appG.Response(http.StatusOK, errMsg.SUCCESS, nil)
}

func GetNotification(c *gin.Context) {
	appG := service.Gin{
		Ctx: c,
	}
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, err)
		return
	}
	res, err := notification_service.Get(int64(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR, err)
		return
	}
	var ns []notification_service.Notification
	for _, v := range res {
		n := notification_service.Notification{}
		json.Unmarshal([]byte(v), &n)
		ns = append(ns, n)
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, ns)
}
