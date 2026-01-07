package umcmp

import (
	"github.com/kweaver-ai/agent-go-common-pkg/cconf"
	"github.com/kweaver-ai/agent-go-common-pkg/src/infra/cmp/icmp"
	// "github.com/kweaver-ai/agent-go-common-pkg/src/infra/config"
)

type Um struct {
	umConf *cconf.UserMgntCfg

	logger icmp.Logger
}

func NewUmCmp(umConf *cconf.UserMgntCfg,
	logger icmp.Logger,
) *Um {
	return &Um{
		umConf: umConf,
		logger: logger,
	}
}
