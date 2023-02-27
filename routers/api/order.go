package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"supply-admin/service"
	errMsg "supply-admin/service/error"
	"supply-admin/service/order_service"
)

type order struct {
	UserID   int64  `json:"userId"`
	MarketID int64  `json:"marketId"`
	OrgID    int    `json:"orgId"`
	Items    string `json:"items"`
	Amount   int    `json:"amount"`
}

func AddOrder(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	var order = &order{}
	err := c.ShouldBindJSON(&order)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, err)
		return
	}

	orderService := order_service.Order{
		UserID:   order.UserID,
		MarketID: order.MarketID,
		OrgID:    order.OrgID,
		Items:    order.Items,
		Amount:   order.Amount,
	}
	id, err := orderService.Add()
	if id == 0 || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.ERROR_CREATE_ORDER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, map[string]int64{
		"order_id": id,
	})
}

func GetOrderForUser(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	userID, ok := c.Params.Get("user_id")
	id, err := strconv.Atoi(userID)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	err, orders := order_service.GetForUser(int64(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ORDER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, orders)
}

func GetOrderForMarket(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	marketID, ok := c.Params.Get("market_id")
	id, err := strconv.Atoi(marketID)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	err, orders := order_service.GetForMarket(int64(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ORDER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, orders)
}

func GetOrderByOrg(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	orgID, ok := c.Params.Get("org_id")
	id, err := strconv.Atoi(orgID)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	err, orders := order_service.GetByOrg(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ORDER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, orders)
}

func UpdateStatus(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	orderID, ok := c.Params.Get("order_id")
	status, ok := c.Params.Get("status")
	id, err := strconv.Atoi(orderID)
	orderStatus, err := strconv.Atoi(status)
	statusStr, ok := order_service.StatusMap[orderStatus]
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	orderService := order_service.OrderStatus{
		ID:          int64(id),
		OrderStatus: orderStatus,
	}
	err = orderService.UpdateStatus()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.UPDATE_ORDER_STATUS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, map[string]string{
		"status": statusStr,
	})
}
