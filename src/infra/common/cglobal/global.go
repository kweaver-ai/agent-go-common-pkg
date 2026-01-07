package cglobal

import (
	"github.com/kweaver-ai/agent-go-common-pkg/cconf"
	"github.com/kweaver-ai/proton-rds-sdk-go/sqlx"
)

var (
	GConfig *cconf.Config // 全局配置
	GDB     *sqlx.DB      // 全局 DB
)
