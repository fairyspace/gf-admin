package middleware

import (
	"context"
	"gf-admin/internal/consts"
	"log"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(r *ghttp.Request) {
	var tokenString = r.Header.Get("Authorization")

	// 验证token
	token, err := jwt.ParseWithClaims(tokenString, &consts.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})

	if err != nil || !token.Valid {
		log.Println("token验证失败", err)
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}

	// 将用户信息存储到上下文中
	if claims, ok := token.Claims.(*consts.JwtClaims); ok {
		ctx := r.GetCtx()
		ctx = context.WithValue(ctx, consts.CtxUserId, claims.Id)
		ctx = context.WithValue(ctx, consts.CtxUserName, claims.UserName)
		r.SetCtx(ctx)
	}

	r.Middleware.Next()
}
