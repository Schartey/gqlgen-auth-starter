package main

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/graphql"
	"github.com/schartey/gqlgen-auth-starter/graphql/resolvers"
	"github.com/spf13/viper"
	"io"
	"syscall"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/handler"
	log "github.com/sirupsen/logrus"

	h "github.com/schartey/gqlgen-auth-starter/graphql/handler"

	nested "github.com/antonfisher/nested-logrus-formatter"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()

	config, err := setupConfig()
	if err != nil {
		log.Panicf("Could not load configs: %s", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize Logging to file and stdout
	initializeLogging(config)

	// Create channel to handle incoming sigint and sigterm
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Add repositories to the resolver so we can store data in resolvers
	keycloakConfig := config.Sub("keycloak")
	resolver := WireUp(ctx, keycloakConfig)

	// Setup server
	server := setupServer(ctx, port, resolver)

	// When we receive sigint or siterm we continue to stopping the server
	<-done
	log.Info("Server Stopped")

	// The server gets 5 seconds time until shutdown
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Debug("Server Exited Properly")
}

func setupConfig() (*viper.Viper, error) {
	configFile := os.Getenv("CONFIG_FILE")
	configPath := os.Getenv("CONFIG_PATH")

	return LoadConfig(configFile, configPath)
}

func initializeLogging(config *viper.Viper) {
	logConfig := config.Sub("log")
	logFile := logConfig.GetString("log-file")

	var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(io.MultiWriter(os.Stdout, file))
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
}

func setupServer(ctx context.Context, port string, resolver *resolvers.RootResolver) *http.Server {
	m := http.NewServeMux()
	server := http.Server{Addr: ":" + port, Handler: m}

	m.Handle("/", handler.Playground("GraphQL playground", "/query"))
	m.Handle("/query", h.AddContext(ctx, handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))))

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	return &server
}
