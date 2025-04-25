package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	UserName string `p:"userName" json:"userName" v:"required#用户名字不能为空"`
	Password string `p:"password" json:"password" v:"required#密码不能为空"`
}

type LoginRes struct {
	Token string `json:"token" dc:"在需要鉴权的接口中header加入Authorization: token"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type LogoutRes struct {
}
