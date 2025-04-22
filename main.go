package main

import (
	_ "gf-admin/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // 导入 MySQL 驱动

	"github.com/gogf/gf/v2/os/gctx"

	"gf-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
