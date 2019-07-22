package domains

import (
	"../utils"
	"../services"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID               int64  `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	SiteID           string `json:"site_id"`
	Permalink        string `json:"permalink"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
				Negative int `json:"negative"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	BuyerReputation struct {
		Tags []interface{} `json:"tags"`
	} `json:"buyer_reputation"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
}

func (user *User) Get() *utils.ApiError {

	if user.ID == 0 {
		return &utils.ApiError{
			Message: "El id de usuario es 0",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%d", utils.UrlUserDev, user.ID)
//	url := fmt.Sprintf("%s%d", utils.UrlUserProd, user.ID)
	data, err := services.GetWithCBreaker(url)
	if err != nil {
		return &utils.ApiError{
			Message: "error del ioutil "+err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if err := json.Unmarshal(data, &user); err != nil {
		return &utils.ApiError{
			Message: "error unmarshall"+err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}


