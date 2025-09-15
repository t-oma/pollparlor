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
	"pollparlor/internal/config"
	"pollparlor/internal/domain"
	"pollparlor/internal/http/handler"
	"pollparlor/internal/http/router"
	"pollparlor/internal/logger"
	pollrepo "pollparlor/internal/repository/poll"
	pollsvc "pollparlor/internal/service/poll"
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

	repo := pollrepo.NewMemoryRepo(seedPolls())
	service := pollsvc.New(repo)
	h := handler.NewPollHandler(service)
	r := router.New(h)

	srv := &http.Server{
		Addr:         cfg.App.Addr + ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("listen")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("shutdown")
	}
}

func seedPolls() []domain.Poll {
	return []domain.Poll{
		{
			ID:    uuid.NewString(),
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
			ID:    uuid.NewString(),
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
