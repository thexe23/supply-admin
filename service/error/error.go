package errMsg

const (
	SUCCESS        = 200
	REDIRECT       = 300
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 10001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10002
	ERROR_AUTH_TOKEN               = 10003
	ERROR_AUTH                     = 10004

	ERROR_REGISTER_FAIL = 20001

	ERROR_ADD_ITEM_FAIL = 30001

	ERROR_GET_USER_FAIL = 40001

	ERROR_GET_ITEM_FAIL = 50001
	UPDATE_ITEM_FAIL    = 50002

	ERROR_GET_ORDER_FAIL     = 60001
	UPDATE_ORDER_STATUS_FAIL = 60002
	ERROR_CREATE_ORDER_FAIL  = 60003
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	REDIRECT:       "重定向",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "用户名或密码错误",

	ERROR_REGISTER_FAIL: "注册账号失败",

	ERROR_ADD_ITEM_FAIL: "添加商品失败",

	ERROR_GET_USER_FAIL: "获取用户失败",

	ERROR_GET_ITEM_FAIL: "获取商品失败",
	UPDATE_ITEM_FAIL:    "更新商品失败",

	ERROR_GET_ORDER_FAIL:     "获取订单失败",
	UPDATE_ORDER_STATUS_FAIL: "变更订单状态失败",
	ERROR_CREATE_ORDER_FAIL:  "创建订单失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
