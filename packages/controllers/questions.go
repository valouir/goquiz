package controllers

import (
	"net/http"

	"github.com/valouir/goquiz/packages/data"

	"github.com/gin-gonic/gin"
)

func GetQuestions(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, data.Questions)
}
