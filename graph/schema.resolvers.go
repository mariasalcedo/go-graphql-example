package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mariasalcedo/go-graphql-example/communication/apimodel"
	"github.com/mariasalcedo/go-graphql-example/communication/client"
	log "github.com/sirupsen/logrus"

	"github.com/mariasalcedo/go-graphql-example/graph/model"
)

// CreateWindFarm is the resolver for the createWindFarm field.
func (r *mutationResolver) CreateWindFarm(ctx context.Context, input *model.NewWindFarm) (*model.WindFarm, error) {
	windFarm := &model.WindFarm{
		Name:      input.Name,
		ID:        uuid.New().String(),
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	}
	r.windFarms = append(r.windFarms, windFarm)
	return windFarm, nil
}

// ID is the resolver for the id field.
func (r *queryResolver) ID(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// WeatherForecasts is the resolver for the weatherForecasts field.
func (r *windFarmResolver) WeatherForecasts(ctx context.Context, obj *model.WindFarm, forecastDays *int) ([]*model.WeatherForecast, error) {
	request := apimodel.ForecastRequest{
		Latitude:     obj.Latitude,
		Longitude:    obj.Longitude,
		ForecastDays: *forecastDays,
	}
	response, err := client.ReadForecast(r.Config, request)
	if err != nil {
		log.WithError(err).Error("Could not obtain forecast for lat: %f lon: %f", obj.Latitude, obj.Longitude)
	}

	weatherForecasts := make([]*model.WeatherForecast, 0)
	for i := range response.Hourly.Time {
		weatherForecasts = append(weatherForecasts, &model.WeatherForecast{
			Time:          response.Hourly.Time[i],
			Temperature:   response.Hourly.Temperature2M[i],
			Precipitation: response.Hourly.Precipitation[i],
			WindSpeed:     response.Hourly.WindSpeed10M[i],
			WindDirection: float64(response.Hourly.WindDirection10M[i]),
		})
	}

	log.Infof("WeatherForecast length %d", len(weatherForecasts))

	return weatherForecasts, err
}

// HasPrecipitationToday is the resolver for the hasPrecipitationToday field.
func (r *windFarmResolver) HasPrecipitationToday(ctx context.Context, obj *model.WindFarm) (bool, error) {
	panic(fmt.Errorf("not implemented: HasPrecipitationToday - hasPrecipitationToday"))
}

// Elevation is the resolver for the elevation field.
func (r *windFarmResolver) Elevation(ctx context.Context, obj *model.WindFarm) (float64, error) {
	request := apimodel.ElevationRequest{
		Latitude:  obj.Latitude,
		Longitude: obj.Longitude,
	}
	response, err := client.ReadElevation(r.Config, request)
	if err != nil {
		log.WithError(err).Error("Could not obtain elevation for lat: %f lon: %f", obj.Latitude, obj.Longitude)
	}
	return response.Elevation[0], err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// WindFarm returns WindFarmResolver implementation.
func (r *Resolver) WindFarm() WindFarmResolver { return &windFarmResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type windFarmResolver struct{ *Resolver }
