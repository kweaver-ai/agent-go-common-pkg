package ginrequestlogger

import (
	"bytes"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/kweaver-ai/agent-go-common-pkg/src/infra/common/chelper/httprequesthelper"
)

// ResponseBodyWriter 用于捕获响应体的ResponseWriter包装器
type ResponseBodyWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

// RequestLogger 请求日志记录器
type RequestLogger struct {
	logger *httprequesthelper.Logger
}

var (
	defaultRequestLogger     *RequestLogger
	defaultRequestLoggerOnce sync.Once
)
