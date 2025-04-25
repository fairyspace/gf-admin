package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf-admin/internal/controller/device_menu"
	"gf-admin/internal/controller/device_role"
	"gf-admin/internal/controller/device_user"
	"gf-admin/internal/controller/login"
	"gf-admin/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/admin", func(group *ghttp.RouterGroup) {
				// 添加自定义响应中间件
				group.Middleware(middleware.CustomResponse)

				// 用户管理相关接口
				group.Group("/acl", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.Group("/user", func(group *ghttp.RouterGroup) {
						group.Bind(
							device_user.NewV1(),
						)
					})
					group.Group("/role", func(group *ghttp.RouterGroup) {
						group.Bind(
							device_role.NewV1(),
						)
					})

					group.Group("/menu", func(group *ghttp.RouterGroup) {
						group.Bind(
							device_menu.NewV1(),
						)
					})
				})

				// 认证相关接口
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.POST("/login", login.NewV1().Login)

					// 需要鉴权的路由组
					group.Middleware(middleware.Auth)
					group.Bind(
						login.NewV1().Logout,
						login.NewV1().Info,
					)

				})
			})
			s.Run()
			return nil
		},
	}
)
