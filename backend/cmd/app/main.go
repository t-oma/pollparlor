// Package main provides a main entry point for the application
package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"pollparlor/internal/bootstrap"
	"pollparlor/internal/config"
	"pollparlor/internal/domain"
	"pollparlor/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	lg := logger.New(logger.Options{
		Level:  cfg.Log.Level,
		Format: cfg.Log.Format,
		Out:    os.Stdout,
	})
	log.Logger = lg

	app, err := bootstrap.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("bootstrap")
	}

	go func() {
		if err := app.Server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("listen")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("shutdown")
	}
}

func seedPolls() []domain.Poll {
	return []domain.Poll{
		{
			ID:    bson.NewObjectID(),
			Title: "Programming languages",
			Author: domain.User{
				UUID:              uuid.NewString(),
				Email:             "user1@example.com",
				CreatedAt:         time.Now(),
				PasswordChangedAt: time.Now(),
			},
			Likes:     0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:    bson.NewObjectID(),
			Title: "Web frameworks",
			Author: domain.User{
				UUID:              uuid.NewString(),
				Email:             "user2@example.com",
				CreatedAt:         time.Now(),
				PasswordChangedAt: time.Now(),
			},
			Likes:     19,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now().Add(time.Hour * 2),
		},
	}
}
