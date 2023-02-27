package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"supply-admin/service"
	errMsg "supply-admin/service/error"
	"supply-admin/service/item_service"
)

type item struct {
	MarketID int64 `json:"marketId"`
	Name     string `json:"name"`
	Price    int `json:"price"`
	Stock    int `json:"stock"`
	Category int `json:"category"`
	ImgUrl   string `json:"imgUrl"`
}

func AddItem(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	var item = &item{}
	err := c.BindJSON(&item)
	if err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, err)
		return
	}

	itemService := item_service.Item{
		MarketID: item.MarketID,
		Name:     item.Name,
		Price:    item.Price,
		Stock:    item.Stock,
		Category: item.Category,
		ImgUrl:   item.ImgUrl,
	}
	id, err := itemService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_ADD_ITEM_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errMsg.SUCCESS, map[string]string{
		"item_id": strconv.FormatInt(id, 10),
	})
}

func UpdateItem(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	itemID, ok := c.Params.Get("item_id")
	item := &item{}
	err := c.ShouldBindJSON(&item)
	id, err := strconv.Atoi(itemID)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}

	itemService := item_service.Item{
		ID:       int64(id),
		MarketID: item.MarketID,
		Name:     item.Name,
		Price:    item.Price,
		Stock:    item.Stock,
		Category: item.Category,
		ImgUrl:   item.ImgUrl,
	}
	err = itemService.Update()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.UPDATE_ORDER_STATUS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, nil)
}

func ToggleItem(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	itemID, ok := c.Params.Get("item_id")
	status, ok := c.Params.Get("on_sale")
	id, err := strconv.Atoi(itemID)
	onSale, err := strconv.ParseBool(status)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}

	itemService := item_service.Item{
		ID:       int64(id),
		OnSale:   onSale,
	}
	err = itemService.Toggle()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.UPDATE_ORDER_STATUS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, nil)
}


func GetItem(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	marketID, ok := c.Params.Get("market_id")
	id, err := strconv.Atoi(marketID)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	err, items := item_service.GetByID(int64(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, items)
}

func GetItemOnSale(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	marketID, ok := c.Params.Get("market_id")
	id, err := strconv.Atoi(marketID)
	if !ok || err != nil {
		appG.Response(http.StatusBadRequest, errMsg.INVALID_PARAMS, nil)
		return
	}
	err, items := item_service.GetOnSale(int64(id))
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, items)
}

func GetAllItem(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	err, items := item_service.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, items)
}
