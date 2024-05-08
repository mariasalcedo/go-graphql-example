package client

import (
	"encoding/json"
	"fmt"
	"github.com/mariasalcedo/go-graphql-example/communication/apimodel"
	"github.com/mariasalcedo/go-graphql-example/pkg/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestElevationClient(t *testing.T) {
	var c config.Config
	var request apimodel.ElevationRequest

	expected := "{\"elevation\":[38.0]}"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	c = config.Config{
		BaseURL:      svr.URL,
		ElevationURL: "",
		ForceHttp:    true,
	}

	request = apimodel.ElevationRequest{
		Latitude:  52.52,
		Longitude: 13.41,
	}

	defer svr.Close()
	res, err := ReadElevation(c, request)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	var expectedFmt apimodel.ElevationResponse
	if err := json.Unmarshal([]byte(expected), &expectedFmt); err != nil {
		log.WithError(err).Error(">>> Error unmarshalling response")
	}

	if expectedFmt.Elevation[0] != res.Elevation[0] {
		t.Errorf("expected res to be %f got %f", expectedFmt.Elevation[0], res.Elevation[0])
	}
}
