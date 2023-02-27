package redis

import (
	. "github.com/go-redis/redis/v8"
	"supply-admin/conf"
)

var Rdb *Client

func Setup() {
	Rdb = NewClient(&Options{
		Addr:     conf.RedisSetting.Host,
		Password: conf.RedisSetting.Password,
	})
}
