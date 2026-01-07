package chttpinject

import (
	"sync"

	"github.com/kweaver-ai/agent-go-common-pkg/src/drivenadapter/httpaccess/umhttpaccess"
	"github.com/kweaver-ai/agent-go-common-pkg/src/port/driven/ihttpaccess/iumacc"
	"devops.aishu.cn/AISHUDevOps/DIP/_git/mdl-go-lib/logger"
)

var (
	umOnce sync.Once
	umImpl iumacc.UmHttpAcc
)

func NewUmHttpAcc() iumacc.UmHttpAcc {
	umOnce.Do(func() {
		umImpl = umhttpaccess.NewUmHttpAcc(
			logger.GetLogger(),
		)
	})

	return umImpl
}
