package base

import (
	"ecom-mono-go/utils/apperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewBaseHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleError(ctx *gin.Context, err error) {
	if errResp, ok := err.(*apperror.AppError); ok {
		errcode := errResp.Code 
		if errcode == http.StatusNotFound {
			errcode = http.StatusBadRequest
		}
		ctx.JSON(int(errcode),gin.H{"err":errResp})
		return
	}
	ctx.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
}