package domains

import (
	"../utils"
	"../services"
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	ID               string    `json:"id"`
	Name         string `json:"name"`
	Locale        string `json:"locale"`
	CurrencyId           string      `json:"currency_id"`
	DecimalSeparator        string      `json:"decimal_separator"`
	ThousandsSeparator        string      `json:"thousands_separator"`
	TimeZone        string      `json:"time_zone"`
	GeoInformation struct {
		Location struct{
			Latitude	float64 `json:"latitude"`
			Longitude	float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		Id           interface{} `json:"id"`
		Name           interface{} `json:"name"`
	} `json:"states"`

}

func (country *Country) Get() *utils.ApiError{

	if country.ID == "" {
		return &utils.ApiError{
			Message: "El id de pais es vacio",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountryDev, country.ID)
//	url := fmt.Sprintf("%s%s", utils.UrlCountryProd, country.ID)

	data, err := services.GetWithCBreaker(url)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil{
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}
