package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLimitOffset(ctx *gin.Context) (intLimit int, offset int) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ = strconv.Atoi(limit)
	offset = (intPage - 1) * intLimit

	return
}
