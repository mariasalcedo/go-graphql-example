package client

import (
	"encoding/json"
	"fmt"
	"github.com/mariasalcedo/go-graphql-example/communication/apimodel"
	"github.com/mariasalcedo/go-graphql-example/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForecastClient(t *testing.T) {
	var c config.Config
	var request apimodel.ForecastRequest

	expected := `
{
  "latitude": 52.52,
  "longitude": 13.419998,
  "generationtime_ms": 0.03409385681152344,
  "utc_offset_seconds": 0,
  "timezone": "GMT",
  "timezone_abbreviation": "GMT",
  "elevation": 38.0,
  "current_units": {
    "time": "iso8601",
    "interval": "seconds",
    "precipitation": "mm"
  },
  "current": {
    "time": "2023-12-07T01:30",
    "interval": 900,
    "precipitation": 0.00
  },
  "hourly_units": {
    "time": "iso8601",
    "temperature_2m": "°C",
    "precipitation": "mm",
    "wind_speed_10m": "km/h",
    "wind_direction_10m": "°"
  },
  "hourly": {
    "time": [
      "2023-12-07T00:00",
      "2023-12-07T01:00",
      "2023-12-07T02:00",
      "2023-12-07T03:00",
      "2023-12-07T04:00",
      "2023-12-07T05:00",
      "2023-12-07T06:00",
      "2023-12-07T07:00",
      "2023-12-07T08:00",
      "2023-12-07T09:00",
      "2023-12-07T10:00",
      "2023-12-07T11:00",
      "2023-12-07T12:00",
      "2023-12-07T13:00",
      "2023-12-07T14:00",
      "2023-12-07T15:00",
      "2023-12-07T16:00",
      "2023-12-07T17:00",
      "2023-12-07T18:00",
      "2023-12-07T19:00",
      "2023-12-07T20:00",
      "2023-12-07T21:00",
      "2023-12-07T22:00",
      "2023-12-07T23:00"
    ],
    "temperature_2m": [
      0.3,
      0.3,
      0.3,
      0.3,
      0.3,
      0.3,
      0.2,
      0.2,
      0.3,
      0.4,
      0.5,
      0.9,
      1.1,
      1.2,
      1.0,
      0.8,
      0.6,
      0.4,
      0.2,
      0.2,
      0.2,
      0.3,
      0.2,
      0.2
    ],
    "precipitation": [
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.10,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00,
      0.00
    ],
    "wind_speed_10m": [
      4.3,
      4.2,
      4.3,
      4.3,
      3.1,
      3.1,
      2.9,
      2.3,
      2.6,
      2.8,
      3.4,
      3.2,
      3.2,
      4.2,
      1.1,
      1.9,
      2.3,
      3.6,
      4.6,
      3.7,
      3.8,
      5.7,
      6.9,
      4.4
    ],
    "wind_direction_10m": [
      222,
      200,
      204,
      228,
      216,
      225,
      210,
      231,
      236,
      230,
      238,
      243,
      270,
      290,
      270,
      158,
      198,
      143,
      129,
      151,
      163,
      125,
      118,
      125
    ]
  }
}
`
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	c = config.Config{
		BaseURL:      svr.URL,
		ElevationURL: "",
		ForecastURL:  "",
		ForceHttp:    true,
	}

	request = apimodel.ForecastRequest{
		Latitude:     52.52,
		Longitude:    13.41,
		ForecastDays: 1,
	}

	defer svr.Close()
	res, err := ReadForecast(c, request)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	var expectedFmt apimodel.ForecastResponse
	if err := json.Unmarshal([]byte(expected), &expectedFmt); err != nil {
		log.WithError(err).Error(">>> Error unmarshalling response")
	}

	if assert.NotNil(t, res) {
		assert.Equal(t, expectedFmt, res)
	}
}
