package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf-admin/internal/controller/device_user"
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
				group.Group("/acl", func(group *ghttp.RouterGroup) {
					group.Group("/user", func(group *ghttp.RouterGroup) {
						group.Bind(
							device_user.NewV1(),
						)
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
