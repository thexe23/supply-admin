package routers

import (
	"github.com/gin-gonic/gin"
	"supply-admin/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("api/auth/login", api.GetAuth)
	r.POST("api/auth/register", api.Register)
	r.GET("api/loading", api.Loading)
	guardApi := r.Group("api/guard")/*.Use(middleware.AuthChecker())*/
	{
		guardApi.GET("user", api.GetUsers)
		guardApi.GET("user/:user_id", api.GetUserByID)
		guardApi.GET("item/", api.GetAllItem)
		guardApi.GET("item/:market_id", api.GetItem)
		guardApi.GET("item/:market_id/on_sale", api.GetItemOnSale)
		guardApi.GET("order/user/:user_id", api.GetOrderForUser)
		guardApi.GET("order/market/:market_id", api.GetOrderForMarket)
		guardApi.GET("order/org/:org_id", api.GetOrderByOrg)
		guardApi.GET("transport", api.GetAllTransports)
		guardApi.GET("transport/source/:source_id", api.GetTransportsForSource)
		guardApi.GET("transport/target/:target_id", api.GetTransportsForTarget)
		guardApi.GET("notification/:id", api.GetNotification)
		guardApi.POST("oss/upload", api.GetUploadToken)
		guardApi.POST("item", api.AddItem)
		guardApi.POST("order", api.AddOrder)
		guardApi.POST("transport", api.AddTransport)
		guardApi.POST("notification", api.AddNotification)
		guardApi.PUT("item/:item_id", api.UpdateItem)
		guardApi.PUT("order/:order_id/:status", api.UpdateStatus)
		guardApi.PUT("user/:user_id", api.UpdateUser)
		guardApi.PUT("transport/:id/:status", api.UpdateTransportStatus)
		guardApi.DELETE("user/:user_id", api.DeleteUser)
		guardApi.DELETE("item/:item_id/:on_sale", api.ToggleItem)
	}

	return r
}
