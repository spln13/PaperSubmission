package middleware

import (
	"PaperSubmission/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var jwtKey = []byte("key_spln")

type Claims struct {
	UserId int64
	jwt.RegisteredClaims
}

// ReleaseToken 颁发管理员专属token
// UserType -> 用户等级标识; 1 -> 学生; 2 -> 教师; 3 -> 管理员
func ReleaseToken(ID int64) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			//ExpiresAt: expirationTime.Unix(),
			IssuedAt: &jwt.NumericDate{Time: time.Now()},
			Issuer:   "linan",
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token != nil {
		if key, ok := token.Claims.(*Claims); ok {
			if token.Valid {
				return key, true
			} else {
				return key, false
			}
		}
	}
	return nil, false
}

// UserJWTMiddleware 判断用户token是否合法，解析token获取user身份
func UserJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.NewCommonResponse(402, "token不存在"))
			c.Abort()
			return
		}
		fmt.Println(tokenStr)
		//验证token
		tokenStruck, ok := ParseToken(tokenStr)
		if !ok {
			c.JSON(http.StatusBadRequest, utils.NewCommonResponse(403, "token不正确"))
			c.Abort() //阻止执行
			return
		}
		//token超时
		fmt.Println(tokenStruck.UserId)
		if time.Now().Unix() > tokenStruck.ExpiresAt.Time.Unix() {
			// token超时, 清空token
			c.SetCookie("token", "", -1, "/", "localhost:8080", true, false)
			c.SetCookie("username", "", -1, "/", "localhost:8080", true, false)
			c.JSON(http.StatusBadRequest, utils.NewCommonResponse(402, "token过期"))
			c.Abort() //阻止执行
			return
		}
		c.Set("userID", tokenStruck.UserId)
		c.Next()
	}
}
