package apimodel

type ElevationRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ElevationResponse struct {
	Elevation []float64 `json:"elevation"`
}
