package chttpinject

import (
	"sync"

	"github.com/kweaver-ai/agent-go-common-pkg/src/drivenadapter/httpaccess/bizdomainhttp"
	"github.com/kweaver-ai/agent-go-common-pkg/src/port/driven/ihttpaccess/ibizdomainacc"
	"github.com/kweaver-ai/kweaver-go-lib/logger"
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
