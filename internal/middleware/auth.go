package middleware

import (
	"gf-admin/internal/consts"
	"log"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(r *ghttp.Request) {
	var tokenString = r.Header.Get("Authorization")

	// 验证token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})
	if err != nil || !token.Valid {
		log.Println("token验证失败", err)
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	r.Middleware.Next()
}
