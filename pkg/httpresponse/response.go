package httpresponse

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/pkg/errors"
	"github.com/xiaodulala/admin-layout/pkg/log"
	"net/http"
)

// ErrResponse 用来返回错误消息的结构
type ErrResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
}

// WriteResponse 解析错误中的错误码等消息，通过gin框架返回
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
