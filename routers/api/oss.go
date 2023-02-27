package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"supply-admin/oss"
	"supply-admin/service"
	errMsg "supply-admin/service/error"
)

type OssResp struct {
	Hash string `json:"hash"`
	Key string `json:"key"`
}

func GetUploadToken(c *gin.Context) {
	appG := service.Gin{Ctx: c}
	token := oss.UploadToken()
	mr, err := c.Request.MultipartReader()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	form, err := mr.ReadForm(128)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	var name string
	var data []byte
	for _, v := range form.File {
		name = v[0].Filename
		f, err := v[0].Open()
		if err != nil {
			appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
			return
		}
		data, err = ioutil.ReadAll(f)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
			return
		}
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	_, err = part.Write(data)
	//将额外参数也写入到multipart
	writer.WriteField("token", token)
	err = writer.Close()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	req, err := http.NewRequest("POST", "http://upload.qiniup.com/", body)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	if resp.StatusCode != 200 {
		appG.Response(resp.StatusCode, errMsg.ERROR, nil)
		return
	}
	ossResp := OssResp{}
	respBody, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &ossResp)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errMsg.ERROR_GET_ITEM_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errMsg.SUCCESS, ossResp.Key)
}
