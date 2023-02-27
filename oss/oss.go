package oss

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"supply-admin/conf"
)

func UploadToken() string{
	mac := qbox.NewMac(conf.OssSetting.AccessKey, conf.OssSetting.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: conf.OssSetting.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	return upToken
}


