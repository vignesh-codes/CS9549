package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type CustomLogFormatterParams struct {
	Request     *http.Request
	Response    string
	Header      http.Header
	RequestBody string

	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
	// isTerm shows whether does gin's output descriptor refers to a terminal.
	isTerm bool
	// BodySize is the size of the Response Body
	BodySize int
	// Keys are the keys set on the request's context.
	Keys map[string]interface{}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func GetRequestAndResponseLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		var req_body, res_body map[string]interface{}
		_ = json.Unmarshal(blw.body.Bytes(), &res_body)
		_ = json.Unmarshal(body, &req_body)

		param := CustomLogFormatterParams{
			Request: c.Request,
		}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = c.ClientIP()
		param.StatusCode = c.Writer.Status()
		param.Path = c.Request.RequestURI
		param.BodySize = c.Writer.Size()
		param.Method = c.Request.Method
		param.Header = c.Request.Header
		param.RequestBody = fmt.Sprintf("%s", req_body)
		if param.Request.Method == "POST" {
			param.Response = fmt.Sprintf("%s", res_body)
		} else {
			param.Response = ""
		}

		strLog := fmt.Sprintf("{\"client_ip\": \"%s\", \"timestamp\": \"%s\", \"method\": \"%s\", \"path\": \"%s\", \"req_proto\": \"%s\", \"status_code\": \"%d\", \"latency\": \"%d\", \"req_user_agent\": \"%s\", \"error_message\": \"%s\", \"device_id\": \"%s\", \"user_uid\": \"%s\",  \"request_body\": \"%s\", \"response_body\": \"%s\"}\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			int(param.Latency.Milliseconds()),
			param.Request.UserAgent(),
			param.ErrorMessage,
			getDeviceId(param),
			getUserUid(param),
			param.RequestBody,
			param.Response,
		)
		fmt.Fprint(gin.DefaultWriter, strLog)
	}
}

func getDeviceId(param CustomLogFormatterParams) string {
	return "internal-service"
}

func getUserUid(param CustomLogFormatterParams) string {
	return "internal-service"
}
