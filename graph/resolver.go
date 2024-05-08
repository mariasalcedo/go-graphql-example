package graph

import (
	"github.com/mariasalcedo/go-graphql-example/graph/model"
	"github.com/mariasalcedo/go-graphql-example/pkg/config"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Config    config.Config
	windFarms []*model.WindFarm
}
