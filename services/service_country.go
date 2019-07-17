package services
import	"../domains"
import "../utils"


func	GetCountryFromApi(countryID string)(*domains.Country,*utils.ApiError){
	country:= &domains.Country{
		ID:countryID,
	}
	if err:=country.Get(); err!= nil {
		return nil, err
	}
	return country,nil
}
