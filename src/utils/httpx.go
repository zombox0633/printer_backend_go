package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrInvalidParam = errors.New("invalid param")

func ParamUint(ctx *gin.Context, name string) (uint, bool) {
	idStr := ctx.Param(name)
	u64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || u64 == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid " + name + " 😿"})
		return 0, false
	}
	return uint(u64), true
}

func ParamInt64(ctx *gin.Context, name string) (int64, bool) {
	idStr := ctx.Param(name)
	i64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || i64 <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid " + name + " 😿"})
		return 0, false
	}
	return i64, true
}
