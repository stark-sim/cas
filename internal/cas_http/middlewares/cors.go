package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
CORS 解决跨域中间件
*/
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		//origin := c.Request.Header.Get("origin")
		//if len(origin) == 0 {
		//	origin = c.Request.Header.Get("Origin")
		//}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Token, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		// 不是每一个请求都要返回 json
		//c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		logrus.Printf("[CORS] set cors success")
		c.Next()
	}
}
