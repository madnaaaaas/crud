package logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/madnaaaaas/crud/pkg/utils"
	"go.uber.org/zap"
	"io"
)

type logPayload struct {
	Method   string      `json:"method"`
	Url      string      `json:"url"`
	Status   int         `json:"status"`
	Errors   interface{} `json:"errors,omitempty"`
	Ip       string      `json:"ip"`
	Request  interface{} `json:"request,omitempty"`
	Response interface{} `json:"response,omitempty"`
}

type LogMiddleware struct {
	log *zap.Logger
}

func NewLogMiddleware(log *zap.Logger) *LogMiddleware {
	return &LogMiddleware{log: log}
}

func (lm *LogMiddleware) Logging(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := io.ReadAll(tee)
	c.Request.Body = io.NopCloser(&buf)

	c.Next()

	payload := logPayload{
		Method: c.Request.Method,
		Url:    c.Request.RequestURI,
		Status: c.Writer.Status(),
		Ip:     c.ClientIP(),
	}

	if utils.IsHTTPStatusSuccess(payload.Status) {
		payload.Request = string(body)
		resp, ok := utils.GetResponse(c)
		if ok {
			payload.Response = resp
		}
	}

	if len(c.Errors) > 0 {
		payload.Errors = c.Errors.String()
		lm.log.Error("error", zap.Any("info", payload))
	} else {
		lm.log.Info("success", zap.Any("info", payload))
	}

}
