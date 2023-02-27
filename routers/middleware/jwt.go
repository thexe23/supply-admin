package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	errMsg "supply-admin/service/error"
	"supply-admin/util"
)

func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = errMsg.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = errMsg.ERROR_AUTH
		} else {
			_, err := util.VerifyToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errMsg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = errMsg.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != errMsg.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errMsg.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
