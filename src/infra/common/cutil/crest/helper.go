package crest

import (
	"devops.aishu.cn/AISHUDevOps/DIP/_git/mdl-go-lib/rest"
	"github.com/pkg/errors"
)

func GetRestHttpErr(err error) (httpError *rest.HTTPError, b bool) {
	if err == nil {
		return
	}

	if errors.As(err, &httpError) {
		b = true
		return
	}

	return
}
