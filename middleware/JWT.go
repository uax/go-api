package middleware

import (
	"api/base"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	TokenNotExits       = errors.New("token not exist")
	TokenValidateFailed = errors.New("token validate failed")
	ClaimsKey           = "UniqueClaimsKey"
	SignKey             = "test"
)

// JwtAuth jwt
type JwtAuth struct {
	SignKey []byte
}

// GenerateToken Generate Token
func (jwtAuth JwtAuth) GenerateToken(tokenExpireTime int64 /* 过期时间 */, iss string /* key */) (string, error) {
	now := time.Now().Unix()
	exp := now + tokenExpireTime
	claim := jwt.MapClaims{
		"iss": iss,
		"iat": now,
		"exp": exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(jwtAuth.SignKey)
	return tokenStr, err
}

//ParseToken parse token
func (jwtAuth JwtAuth) ParseToken(token string) (jwt.Claims, error) {
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtAuth.SignKey, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims.Claims == nil || !tokenClaims.Valid {
		return nil, TokenValidateFailed
	}
	return tokenClaims.Claims, nil
}

//validateToken validate token
func validateToken(ctx *gin.Context) error {
	//	get Token
	tokenStr := ctx.GetHeader("Authorization")
	kv := strings.Split(tokenStr, ",")
	if len(kv) != 2 || kv[0] != "Bearer" {
		return TokenNotExits
	}
	jwtAuth := &JwtAuth{SignKey: []byte(SignKey)}
	claims, err := jwtAuth.ParseToken(kv[1])
	if err != nil {
		return err
	}
	ctx.Set(ClaimsKey, claims)
	return nil
}

//JWT jwt middleware

func JWT(ctx *gin.Context) {
	if err := validateToken(ctx); err != nil {
		base.WrapContext(ctx).Error(401, err.Error())
		ctx.Abort()
		return
	}
	ctx.Next()
}
