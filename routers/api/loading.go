package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"supply-admin/service"
	errMsg "supply-admin/service/error"
	"supply-admin/service/user_service"
	"supply-admin/util"
)

func Loading(c *gin.Context) {
	appG := service.Gin{
		Ctx: c,
	}
	token:= c.Request.Header.Get("Authorization")
	if token == "" {
		appG.Response(http.StatusOK, errMsg.REDIRECT, map[string]string{
			"path": "/login",
		})
		return
	}
	tokenContent, err := util.ExtractTokenMetadata(token)
	if err != nil {
		appG.Response(http.StatusOK, errMsg.REDIRECT, map[string]string{
			"path": "/login",
		})
		return
	}
	err, user := user_service.GetUserByID(tokenContent.UserId)
	if err != nil {
		appG.Response(http.StatusOK, errMsg.REDIRECT, map[string]string{
			"path": "/login",
		})
		return
	}
	if user.Role == 1 {
		appG.Response(http.StatusOK, errMsg.REDIRECT, map[string]string{
			"path": "/shopping",
			"id": strconv.FormatInt(user.ID, 10),
			"market_id": strconv.FormatInt(user.MarketID, 10),
			"org_id": strconv.FormatInt(user.OrgID, 10),
		})
		return
	}
	appG.Response(http.StatusOK, errMsg.REDIRECT, map[string]string{
		"path": "/admin",
		"id": strconv.FormatInt(user.ID, 10),
		"org_id": strconv.FormatInt(user.OrgID, 10),
		"role": strconv.FormatInt(int64(user.Role), 10),
	})
}
