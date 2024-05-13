package apimodel

type ForecastRequest struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	ForecastDays int     `json:"forecast_days"`
}

type ForecastResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationTimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentUnits         struct {
		Time          string `json:"time"`
		Interval      string `json:"interval"`
		Precipitation string `json:"precipitation"`
	} `json:"current_units"`
	Current struct {
		Time          string  `json:"time"`
		Interval      int     `json:"interval"`
		Precipitation float64 `json:"precipitation"`
	} `json:"current"`
	HourlyUnits struct {
		Time             string `json:"time"`
		Temperature2M    string `json:"temperature_2m"`
		Precipitation    string `json:"precipitation"`
		WindSpeed10M     string `json:"wind_speed_10m"`
		WindDirection10M string `json:"wind_direction_10m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time             []string  `json:"time"`
		Temperature2M    []float64 `json:"temperature_2m"`
		Precipitation    []float64 `json:"precipitation"`
		WindSpeed10M     []float64 `json:"wind_speed_10m"`
		WindDirection10M []int     `json:"wind_direction_10m"`
	} `json:"hourly"`
}
