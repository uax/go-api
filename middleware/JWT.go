package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

type MyClaims struct {
	UID int64 `json:"uid"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 1

var MySecret = []byte("Noah")

//GenerateToken generate token
func GenerateToken(uid int64) (string, error) {
	c := MyClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Noah",
		},
	}
	//	使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

//ParseToken parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
	//if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
	//	return claims, nil
	//} else {
	//	return nil, err
	//}
	//if ve, ok := err.(*jwt.ValidationError); ok {
	//	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	//		fmt.Println("That's not even a token")
	//	} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	//		// Token is either expired or not active yet
	//		fmt.Println("Timing is everything")
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println("Couldn't handle this token:", err)
	//	}
	//} else {
	//	fmt.Println("Couldn't handle this token:", err)
	//}
	//return nil, errors.New(err.Error())
}

//JWTMiddleware JWT Middleware
func JWTMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "empty authorization",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		u, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("uid", u.UID)
		c.Next()
	}
}
