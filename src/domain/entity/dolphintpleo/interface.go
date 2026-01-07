package dolphintpleo

import "github.com/kweaver-ai/agent-go-common-pkg/src/domain/valueobject/daconfvalobj"

type IDolphinTpl interface {
	LoadFromConfig(config *daconfvalobj.Config)
	ToString() (str string)
}
