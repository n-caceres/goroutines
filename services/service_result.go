package services

import (
	"../domains"
)
import "../utils"

var result domains.Result

//func	GetResultFromApi(user *domains.User, site *domains.Site,country *domains.Country)(*domains.Result){

func	SetUserForResult(id int64)( *domains.Result,*utils.ApiError) {
	user, apiError := GetUserFromApi(id) //Traigo Usuario
	if apiError != nil {
		return nil,apiError
	}
	result:= domains.Result{
User:user,
	}
	return &result,nil
}
func	SetSiteForResult(result *domains.Result)(*domains.Result,*utils.ApiError) {
	site, apiError := GetSiteFromApi(result.User.SiteID) //Traigo Site, con site id del usuario
	if apiError != nil {
		return nil,apiError
	}
	result.Site=site
	return result,nil
}
func	SetCountryForResult(result *domains.Result)(*domains.Result,*utils.ApiError) {
country,apiError:=GetCountryFromApi(result.User.CountryID)//Traigo Country con countryID del usuario
if apiError!= nil {
return nil,apiError
}
result.Country=country
return result,nil
}

