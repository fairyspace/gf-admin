package consts

import "github.com/golang-jwt/jwt/v5"

// JwtClaims JWT claims结构体
type JwtClaims struct {
	Id       uint   `json:"id"`
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

// ContextKey 用于context的key类型
type ContextKey string

const (
	CtxUserId   ContextKey = "userId"
	CtxUserName ContextKey = "userName"
)
