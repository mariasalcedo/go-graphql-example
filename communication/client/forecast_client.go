package client

import (
	"encoding/json"
	"fmt"
	"github.com/mariasalcedo/go-graphql-example/communication/apimodel"
	"github.com/mariasalcedo/go-graphql-example/pkg/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func ReadForecast(c config.Config, request apimodel.ForecastRequest) (apimodel.ForecastResponse, error) {
	return readForecast(&http.Client{
		Timeout: time.Second * 30,
	}, c, request)
}

func readForecast(client *http.Client, c config.Config, request apimodel.ForecastRequest) (apimodel.ForecastResponse, error) {
	var req *http.Request
	var err error
	log.Infof("[ForecastClient] Building initial url...")
	req, err = http.NewRequest("GET", c.BaseURL+c.ElevationURL, nil)
	if err != nil {
		log.WithError(err).Error(">>> Error creating request")
		return apimodel.ForecastResponse{}, err
	}

	q := req.URL.Query()
	q.Add("latitude", fmt.Sprint(request.Latitude))
	q.Add("longitude", fmt.Sprint(request.Longitude))
	q.Add("current", "precipitation")
	q.Add("hourly", "temperature_2m,precipitation,wind_speed_10m,wind_direction_10m")
	q.Add("forecast_days", fmt.Sprint(request.ForecastDays))
	req.URL.RawQuery = q.Encode()

	if c.ForceHttp {
		req.URL.Scheme = "http"
	}

	log.Infof("[ForecastClient] Calling %s", req.URL)

	payload, err := doRequest(client, req)
	if err != nil {
		log.WithError(err).Error(">>> Error executing request")
		return apimodel.ForecastResponse{}, err
	}

	var page apimodel.ForecastResponse
	if err := json.Unmarshal(payload, &page); err != nil {
		log.WithError(err).Error(">>> Error unmarshalling response")
		return apimodel.ForecastResponse{}, err
	}

	return page, nil
}
