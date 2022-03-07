package httpresponse

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/pkg/errors"
	"github.com/xiaodulala/admin-layout/pkg/log"
	"github.com/xiaodulala/admin-layout/pkg/middleware"
	"net/http"
)

// ErrResponse 用来返回错误消息的结构
type ErrResponse struct {
	RequestId string `json:"requestId,omitempty"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
}

// WriteErrResponse 解析错误中的错误码等消息，通过gin框架返回
func WriteErrResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrResponse{
			RequestId: middleware.GetRequestIDFromContext(c),
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})
		c.Abort()
		return
	}

	WriteSuccessResponse(c, data)
}

type Response struct {
	RequestId string      `json:"requestId,omitempty"`
	Data      interface{} `json:"data"`
}

func WriteSuccessResponse(c *gin.Context, data interface{}) {
	res := Response{
		RequestId: middleware.GetRequestIDFromContext(c),
		Data:      data,
	}
	c.JSON(http.StatusOK, res)
	return
}
