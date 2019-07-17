package controllers

import (
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const(
	paramSiteId = "siteID"
)


func	GetSite(ctx *gin.Context){

	siteID:=ctx.Param(paramSiteId)
	site,apiError:=services.GetSiteFromApi(siteID)
	if apiError!= nil {
		ctx.JSON(apiError.Status,apiError.Message)
		return
	}
	ctx.JSON(http.StatusOK,site)
}

