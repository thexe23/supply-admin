package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"supply-admin/service"
	errMsg "supply-admin/service/error"
	"supply-admin/service/transport_service"
)

type AddTransportReq struct {
	SourceID int64     `json:"sourceId"`
	TargetID int64     `json:"targetId"`
	Item     string    `json:"item"`
	Quantity int       `json:"quantity"`
}

func AddTransport(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	var addReq AddTransportReq
	err := c.ShouldBindJSON(&addReq)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	transportService := transport_service.Transport{
		SourceID: addReq.SourceID,
		TargetID: addReq.TargetID,
		Item:     addReq.Item,
		Quantity: addReq.Quantity,
	}
	id, err := transportService.Add()
	if id == 0 || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, map[string]int64{
		"transport_id": id,
	})
}

func GetAllTransports(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	ts := transport_service.Transport{}
	err, transports := ts.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, transports)
}

func GetTransportsForSource(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	id, ok := c.Params.Get("source_id")
	sid, err := strconv.Atoi(id)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	ts := transport_service.Transport{
		SourceID: int64(sid),
	}
	err, transports := ts.GetForSource()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, transports)
}

func GetTransportsForTarget(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	id, ok := c.Params.Get("target_id")
	tid, err := strconv.Atoi(id)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	ts := transport_service.Transport{
		TargetID: int64(tid),
	}
	err, transports := ts.GetForTarget()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, transports)
}

func UpdateTransportStatus(c *gin.Context) {
	appG := service.Gin{
		Ctx: c,
	}
	idStr := c.Param("id")
	statusStr := c.Param("status")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	ts := transport_service.Transport{
		ID: int64(id),
		Status: status,
	}
	err = ts.UpdateStatus()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, nil)
}
