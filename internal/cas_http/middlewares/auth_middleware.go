package middlewares

import (
	"cas/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

const ResponseWriter = "RESPONSE_WRITER"

/*
AuthMiddleware 认证中间件
*/
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取 cookie
		cookie, err := c.Cookie(tools.CookieName)
		cookie, err = url.PathUnescape(cookie)
		logrus.Printf("=======> cookie: %s", cookie)
		if err != nil && err != http.ErrNoCookie {
			c.Abort()
			return
		}
		// 某些情况可以没有 cookie
		if err == http.ErrNoCookie {
			fmt.Printf("%s\n", c.Request.RequestURI)
			c.Next()
		}
		// 验证 JWT， 顺便把 userId 存于上下文
		customClaims, err := tools.ParseToken(cookie)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("userID", customClaims.UserID)
		// 将 token 保存到上下文中，便于发送 GRPC 请求，由于 GRPC 请求 METADATA key 全小写，所以 Authorization 换成 token
		c.Set("token", fmt.Sprintf("%s%s", tools.JWTHeader, cookie))
		c.Next()
	}
}

// InjectableResponseWriter 将 writer 载入到 ctx 中
type InjectableResponseWriter struct {
	http.ResponseWriter
	Cookie *http.Cookie
}

func (i *InjectableResponseWriter) Write(data []byte) (int, error) {
	if i.Cookie != nil {
		http.SetCookie(i.ResponseWriter, i.Cookie)
	}
	return i.ResponseWriter.Write(data)
}

func WriterMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		injectableResponseWriter := InjectableResponseWriter{
			ResponseWriter: c.Writer,
			Cookie:         nil,
		}
		c.Set(ResponseWriter, &injectableResponseWriter)
		c.Next()
	}
}
