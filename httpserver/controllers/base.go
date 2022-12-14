package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanramli/hacktiv8-assignment-2/httpserver/controllers/views"
)

func WriteJsonResponse(ctx *gin.Context, resp *views.Response) {
	ctx.JSON(resp.Status, resp)
}
