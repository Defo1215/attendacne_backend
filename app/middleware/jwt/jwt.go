package jwt

import (
	"attendance/utils"
	"github.com/gin-gonic/gin"
)

// JWT 自定义中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 400 // 无token，无权限访问
		} else {
			// 解析token
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 400 // token不合法
			} else {
				id := claims.ID
				username := claims.Username

				// 将当前请求的user信息保存到请求的上下文c上
				c.Set("id", id)
				c.Set("username", username)
			}
			//else if time.Now().Unix() > claims.ExpiresAt {
			//	code = result.TokenExpired.Code // token已过期
			//}

		}
		if code != 200 {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  "无权限访问",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
