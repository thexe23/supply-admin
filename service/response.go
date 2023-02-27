package service

import (
	"github.com/gin-gonic/gin"
	errMsg "supply-admin/service/error"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.Ctx.JSON(httpCode, Response{
		Code: errCode,
		Msg:  errMsg.GetMsg(errCode),
		Data: data,
	})
	return
}
