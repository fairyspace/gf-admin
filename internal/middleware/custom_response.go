package middleware

import (
	"mime"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
)

var (
	// 流媒体类型
	streamContentType = []string{contentTypeEventStream, contentTypeOctetStream, contentTypeMixedReplace}
)

// CustomResponse 自定义响应中间件
func CustomResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果响应已经写入，则不进行修改
	if r.Response.BufferLength() > 0 || r.Response.Writer.BytesWritten() > 0 {
		return
	}

	// 如果响应类型是流媒体类型，则不进行修改
	mediaType, _, _ := mime.ParseMediaType(r.Response.Header().Get("Content-Type"))
	for _, ct := range streamContentType {
		if mediaType == ct {
			return
		}
	}

	var (
		err  = r.GetError()
		msg  = getMessage(err)
		res  = r.GetHandlerResponse()
		code = getCode(err)
	)

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    res,
	})
}

func getCode(err error) int {
	if err == nil {
		return 200
	}
	return gerror.Code(err).Code()
}

func getMessage(err error) string {
	if err == nil {
		return "success"
	}
	return err.Error()
}
