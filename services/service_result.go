package services

import (
	"../domains"
)
import "../utils"

var result domains.Result


func	SetUserForResult(id int64)( *domains.Result,*utils.ApiError) {
	user, apiError := GetUserFromApi(id)
	if apiError != nil {
		return nil,apiError
	}
	result:= domains.Result{
User:user,
	}
	return &result,nil
}
func	SetSiteForResult(result *domains.Result)(*domains.Result,*utils.ApiError) {
	site, apiError := GetSiteFromApi(result.User.SiteID)
	if apiError != nil {
		return nil,apiError
	}
	result.Site=site
	return result,nil
}
func	SetCountryForResult(result *domains.Result)(*domains.Result,*utils.ApiError) {
country,apiError:=GetCountryFromApi(result.User.CountryID)
if apiError!= nil {
return nil,apiError
}
result.Country=country
return result,nil
}

