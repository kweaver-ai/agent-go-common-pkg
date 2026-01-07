package chelper

import (
	"github.com/kweaver-ai/agent-go-common-pkg/src/infra/cmp/icmp"
	"devops.aishu.cn/AISHUDevOps/DIP/_git/mdl-go-lib/logger"
	"go.uber.org/zap"
)

var simpleStdoutLogger *zap.SugaredLogger

// GetStdoutLogger 不依赖配置文件，直接输出到stdout
func GetStdoutLogger() icmp.Logger {
	if simpleStdoutLogger != nil {
		return simpleStdoutLogger
	}

	simpleStdoutLogger = logger.GetLogger()

	return simpleStdoutLogger
}
