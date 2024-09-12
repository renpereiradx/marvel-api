package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/renpereiradx/marvel-api/database"
	"github.com/renpereiradx/marvel-api/repository"
	"github.com/rs/cors"
)

type Config struct {
	Port        string
	JwtSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JwtSecret == "" {
		return nil, errors.New("jwt is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, router *mux.Router)) {
	binder(b, b.router)

	handler := cors.AllowAll().Handler(b.router)

	repo, err := database.NewPostgresRepo(b.config.DatabaseUrl)
	if err != nil {
		log.Println("error connecting to database ", err)
		return
	}
	repository.SetRepository(repo)

	log.Println("Starting server on Port: ", b.config.Port)
	if err = http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Println("error starting server", err)
	} else {
		log.Println("server stopped")
	}
}
