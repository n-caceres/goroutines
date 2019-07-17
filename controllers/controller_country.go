package controllers

import (
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const(
	paramCountryId = "countryID"
)


func	GetCountry(ctx *gin.Context){

	countryID:=ctx.Param(paramCountryId)

	country,apiError:=services.GetCountryFromApi(countryID)
	if apiError!= nil {
		ctx.JSON(apiError.Status,apiError.Message)
		return
	}
	ctx.JSON(http.StatusOK,country)
}
