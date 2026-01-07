package chttpinject

import (
	"sync"

	"github.com/kweaver-ai/agent-go-common-pkg/src/drivenadapter/httpaccess/bizdomainhttp"
	"github.com/kweaver-ai/agent-go-common-pkg/src/port/driven/ihttpaccess/ibizdomainacc"
	"devops.aishu.cn/AISHUDevOps/DIP/_git/mdl-go-lib/logger"
)

var (
	bizDomainOnce sync.Once
	bizDomainImpl ibizdomainacc.BizDomainHttpAcc
)

func NewBizDomainHttpAcc() ibizdomainacc.BizDomainHttpAcc {
	bizDomainOnce.Do(func() {
		bizDomainImpl = bizdomainhttp.NewBizDomainHttpAcc(
			logger.GetLogger(),
		)
	})

	return bizDomainImpl
}
