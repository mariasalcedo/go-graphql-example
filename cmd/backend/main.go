package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mariasalcedo/go-graphql-example/graph"
	"github.com/mariasalcedo/go-graphql-example/pkg/config"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var c config.Config

func fromEnv(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

func init() {
	fmt.Println("Initializing configuration")

	c = config.Config{}

	c.BaseURL = fromEnv("BASE_URL")
	if c.BaseURL == "" {
		log.Fatal("BASE_URL not configured")
	}

	c.ElevationURL = fromEnv("ELEVATION_URL")
	if c.ElevationURL == "" {
		log.Fatal("ELEVATION_URL not configured")
	}

	c.ForecastURL = fromEnv("FORECAST_URL")
	if c.ForecastURL == "" {
		log.Fatal("FORECAST_URL not configured")
	}

	var err error

	forceHttp := fromEnv("FORCE_HTTP")
	if forceHttp == "" {
		forceHttp = "false"
	}
	c.ForceHttp, err = strconv.ParseBool(forceHttp)
	if err != nil {
		log.Fatalf("FORCE_HTTP '%s' cannot be converted to bool", forceHttp)
	}
}

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting Log Level
	// When using stackdriver: logging does not recognize "level" but needs "severity" instead
	log.SetFormatter(&log.JSONFormatter{FieldMap: log.FieldMap{log.FieldKeyLevel: "severity"}})
	log.SetOutput(os.Stdout)
	logLevel, err := log.ParseLevel(fromEnv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	// Setting healthz livez
	http.HandleFunc("/.well-known/live", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	})
	http.HandleFunc("/.well-known/ready", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	})

	// Setting graphQL server
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					Config: c,
				},
			},
		),
	)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
