package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const(
	paramUserId = "userID"
)


func	GetUser(ctx *gin.Context){

	userID:=ctx.Param(paramUserId)

	id,err:= strconv.ParseInt(userID,10,64)
	if  err!=nil{
		apiError:= &utils.ApiError{
			Message: err.Error(),
			Status:http.StatusBadRequest,
		}
		ctx.JSON(apiError.Status,apiError.Message)
		return
	}
	user,apiError:=services.GetUserFromApi(id)
	if apiError!= nil {
		ctx.JSON(apiError.Status,apiError.Message)
		return
	}
	ctx.JSON(http.StatusOK,user)
}
