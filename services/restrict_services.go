package services

import (
	"github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
	"time"
)
var cb *gobreaker.CircuitBreaker
var rafagaRequest = make (chan int,15)


func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

func GetWithCBreaker(url string) ([]byte, error) {


	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		select {
		case rafagaRequest<-1: // Put 2 in the channel unless it is full
		default:
			time.Sleep(time.Second*15)
			for i:=1;i<=15 ;i++  {
				 <-rafagaRequest
			}

		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}
