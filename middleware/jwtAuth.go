package middleware

import (
	"app/pkg/error_code"
	"app/pkg/response"
	"app/pkg/utils"
	"app/service"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Resp().SetHttpCode(http.StatusUnauthorized).Fail(ctx, error_code.Text(error_code.StatusUnauthorized))
			ctx.Abort() //结束后续操作
			return
		}
		log.Print("token:", authHeader)

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Resp().SetHttpCode(http.StatusUnauthorized).Fail(ctx, error_code.Text(error_code.StatusUnauthorized))
			ctx.Abort()
			return
		}

		//解析token包含的信息
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.Resp().SetHttpCode(http.StatusUnauthorized).Fail(ctx, error_code.Text(error_code.StatusUnauthorized))
			ctx.Abort()
			return
		}

		if err := CheckUserInfo(claims); err != nil {
			response.Resp().SetHttpCode(http.StatusUnauthorized).Fail(ctx, error_code.Text(error_code.StatusUnauthorized))
			ctx.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		ctx.Set("claims", claims)
		ctx.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息

	}
}

//检查用户名信息
func CheckUserInfo(claims *utils.CustomClaims) error {
	username := claims.Name
	password := claims.Password
	userService := service.UserService{}
	user := userService.GetUserNamePassword(username, password)
	if user == nil {
		return errors.New("User not found")
	}
	// userDao := dao.UserDao{}
	// user, err = userDao.(username, password)
	// if err != nil {
	// 	return err
	// }
	//获取数据库用户名及密码
	if username == user.Name && password == user.Password {
		return nil
	}
	return errors.New("用户密码错误")
}
