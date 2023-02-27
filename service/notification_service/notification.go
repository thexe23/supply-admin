package notification_service

import (
	"context"
	"encoding/json"
	redis2 "github.com/go-redis/redis/v8"
	"strconv"
	"supply-admin/redis"
)

type Notification struct {
	SourceID int64  `json:"sourceId"`
	TargetID int64  `json:"targetId"`
	Type     int    `json:"type"`
	Content  string `json:"content"`
}

func Add(msg Notification) error {
	rdb := redis.Rdb
	id := strconv.FormatInt(msg.TargetID, 10)
	content,_ := json.Marshal(msg)
	z := &redis2.Z{
		Member: string(content),
	}
	return rdb.ZAdd(context.Background(), id, z).Err()
}

func Get(id int64) ([]string, error) {
	rdb := redis.Rdb
	return rdb.ZRange(context.Background(), strconv.FormatInt(id, 10), 0, -1).Result()
}
