package dolphintpleo

import (
	"github.com/kweaver-ai/agent-go-common-pkg/src/domain/enum/cdaenum"
)

type DolphinTplEo struct {
	Key   cdaenum.DolphinTplKey `json:"key"`
	Name  string                `json:"name"`
	Value string                `json:"value"`
}
