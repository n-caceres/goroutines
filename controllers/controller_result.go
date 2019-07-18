package controllers
//github.com/stretchr/testify/assert
import (
	"../domains"
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

func	GetResult(ctx *gin.Context){

	c1:=make(chan *utils.ApiError)
	defer close(c1)
	c2:=make(chan *domains.Result)
	defer close(c2)
	userID:=ctx.Param(paramUserId)
	var wg sync.WaitGroup

	id,err:= strconv.ParseInt(userID,10,64)
	if  err!=nil{
		apiError:= &utils.ApiError{
			Message: err.Error(),
			Status:http.StatusBadRequest,
		}
		ctx.JSON(apiError.Status,apiError.Message)
		return
	}

	result,error:=services.SetUserForResult(id)
	if error!=nil{
		ctx.JSON(error.Status,error.Message)
		return
	}
	wg.Add(1)
	go func() {
		result,error=services.SetSiteForResult(result)
		c1<-error
		c2<-result
	}()
	wg.Add(1)
	go func() {
		result,error:=services.SetCountryForResult(result)
		c1<-error
		c2<-result
	}()
	var errores *utils.ApiError
	go func() {
		for e:=range c1 {
			if e!=nil {
				errores=e
			}
		}
	}()
	go func() {
		for e:=range c2 {
			wg.Done()
			if e.Country!=nil {continue}
			if e.Site!=nil {continue}
		}
	}()

	if errores!=nil{
		ctx.JSON(errores.Status,errores.Message)
		return
	}

	//Limpio el waiting group de las corridas correctas.
	wg.Wait()
	ctx.JSON(http.StatusOK,result)
}

