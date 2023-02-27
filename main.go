package main

import (
	"fmt"
	"log"
	"net/http"
	"supply-admin/conf"
	"supply-admin/models"
	"supply-admin/redis"
	"supply-admin/routers"
)

func init() {
	conf.Setup()
	models.Setup()
	redis.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	readTimeout := conf.ServerSetting.ReadTimeout
	writeTimeout := conf.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", conf.ServerSetting.HttpPort)
	maxHeadBytes := 1<<20

	server := &http.Server{
		Addr: endPoint,
		Handler: routersInit,
		ReadTimeout: readTimeout,
		WriteTimeout: writeTimeout,
		MaxHeaderBytes: maxHeadBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
