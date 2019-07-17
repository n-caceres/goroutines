package domains
import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site struct {
	ID               string    `json:"id"`
	Name         string `json:"name"`
	CountryID        string `json:"country_id"`
	SalesFeedMode           string      `json:"sale_fees_mode"`
	MercadoPagoVersion        string      `json:"mercado_pago_version"`
	DefaultCurrencyId        string      `json:"default_currency_id"`
	ImmediatePayment        string      `json:"immediate_payment"`
	PaymentMethodIds		[]string	`json:"payment_methods_id"`
	Settings struct {
		IdentificationTypes           []string `json:"identification_types"`
		TaxPayerTypes           []string `json:"tax_payer_types"`
		IdentificationTypesRules interface{} `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		Id           interface{} `json:"id"`
		Symbol           interface{} `json:"symbol"`
	} `json:"currencies"`
	Categories []struct {
		Id           interface{} `json:"id"`
		Name           interface{} `json:"name"`
	} `json:"settings"`

}

func (site *Site) Get() *utils.ApiError{

	if site.ID == "" {
		return &utils.ApiError{
			Message: "El id de site es vacio",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

	response, err := http.Get(url)
	if err != nil{
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &site); err != nil{
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}
