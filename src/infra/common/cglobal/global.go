package cglobal

import (
	"github.com/kweaver-ai/agent-go-common-pkg/cconf"
	"devops.aishu.cn/AISHUDevOps/ONE-Architecture/_git/proton-rds-sdk-go/sqlx"
)

var (
	GConfig *cconf.Config // 全局配置
	GDB     *sqlx.DB      // 全局 DB
)
