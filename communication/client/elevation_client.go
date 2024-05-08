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

func ReadElevation(c config.Config, request apimodel.ElevationRequest) (apimodel.ElevationResponse, error) {
	return readElevation(&http.Client{
		Timeout: time.Second * 30,
	}, c, request)
}

func readElevation(client *http.Client, c config.Config, request apimodel.ElevationRequest) (apimodel.ElevationResponse, error) {
	var req *http.Request
	var err error
	log.Infof("[ElevationClient] Building initial url...")
	req, err = http.NewRequest("GET", c.BaseURL+c.ElevationURL, nil)
	if err != nil {
		log.WithError(err).Error(">>> Error creating request")
		return apimodel.ElevationResponse{}, err
	}

	q := req.URL.Query()
	q.Add("latitude", fmt.Sprint(request.Latitude))
	q.Add("longitude", fmt.Sprint(request.Longitude))
	req.URL.RawQuery = q.Encode()

	if c.ForceHttp {
		req.URL.Scheme = "http"
	}

	log.Infof("[ElevationClient] Calling %s", req.URL)

	payload, err := doRequest(client, req)
	if err != nil {
		log.WithError(err).Error(">>> Error executing request")
		return apimodel.ElevationResponse{}, err
	}

	var page apimodel.ElevationResponse
	if err := json.Unmarshal(payload, &page); err != nil {
		log.WithError(err).Error(">>> Error unmarshalling response")
		return apimodel.ElevationResponse{}, err
	}

	return page, nil
}
