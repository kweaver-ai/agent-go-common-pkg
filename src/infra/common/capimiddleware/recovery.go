package capimiddleware

import (
	"fmt"
	"log"

	"github.com/kweaver-ai/agent-go-common-pkg/src/infra/common/capierr"
	"github.com/kweaver-ai/agent-go-common-pkg/src/infra/common/chelper/cenvhelper"
	"github.com/kweaver-ai/agent-go-common-pkg/src/infra/common/chelper/panichelper"
	"devops.aishu.cn/AISHUDevOps/DIP/_git/mdl-go-lib/logger"
	"devops.aishu.cn/AISHUDevOps/DIP/_git/mdl-go-lib/rest"
	"github.com/gin-gonic/gin"
)

// Recovery recover中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				_logger := logger.GetLogger()

				// 1、记录日志
				panicLogMsg := panichelper.PanicTraceErrLog(e)
				_logger.Errorln(panicLogMsg)

				if cenvhelper.IsLocalDev() {
					log.Print(panicLogMsg)
				}

				// 2、返回错误信息
				err := capierr.New500Err(c, fmt.Sprintf("internal error: %v", e))
				rest.ReplyError(c, err)

				c.Abort()

				return
			}
		}()

		c.Next()
	}
}
