package login

import (
	"context"
	v1 "gf-admin/api/login/v1"
	"gf-admin/internal/consts"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
)

type Login struct {
}

func (l *Login) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRes, error) {
	var user entity.DeviceUser
	var res v1.LoginRes

	orm := dao.DeviceUser.Ctx(ctx)
	err := orm.Where(do.DeviceUser{
		UserName: req.UserName,
	}).Scan(&user)

	if err != nil {
		return nil, gerror.New("用户名或者密码错误")
	}

	//用户不存在
	if user.Id == 0 {
		return nil, gerror.New("用户不存在")
	}

	//密码错误
	if user.Password != req.Password {
		return nil, gerror.New("密码错误")
	}

	// 生成 JWT token
	claims := consts.JwtClaims{
		Id:       uint(user.Id),
		UserName: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	pre_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	Token, err := pre_token.SignedString([]byte(consts.JwtKey)) // 请替换为实际的密钥
	if err != nil {
		return nil, gerror.New("生成token失败")
	}

	res.Token = Token
	return &res, nil
}

func (l *Login) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutRes, error) {
	return &v1.LogoutRes{}, nil
}

func New() *Login {
	return &Login{}
}
