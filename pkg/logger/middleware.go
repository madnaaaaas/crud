package logger

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type logPayload struct {
	Method   string      `json:"method"`
	Url      string      `json:"url"`
	Status   int         `json:"status"`
	Errors   interface{} `json:"errors,omitempty"`
	Ip       string      `json:"ip"`
	Request  string      `json:"request,omitempty"`
	Response string      `json:"response,omitempty"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

type LogMiddleware struct {
	log *zap.Logger
}

func NewLogMiddleware(log *zap.Logger) *LogMiddleware {
	return &LogMiddleware{log: log}
}

func (lm *LogMiddleware) Logging(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := io.ReadAll(tee)
	c.Request.Body = io.NopCloser(&buf)

	c.Next()

	payload := logPayload{
		Method: c.Request.Method,
		Url:    c.Request.RequestURI,
		Status: c.Writer.Status(),
		Errors: c.Errors.String(),
		Ip:     c.ClientIP(),
	}

	if payload.Status >= http.StatusOK && payload.Status < http.StatusMultipleChoices {
		payload.Request = string(body)
		payload.Response = blw.body.String()
	}

	if len(c.Errors) > 0 {
		lm.log.Error("error", zap.Any("info", payload))
	} else {
		lm.log.Info("success", zap.Any("info", payload))
	}

}
