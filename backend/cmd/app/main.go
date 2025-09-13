// Package main provides a main entry point for the application
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"pollparlor/internal/domain"
	"pollparlor/internal/http/handler"
	"pollparlor/internal/http/router"
	pollrepo "pollparlor/internal/repository/poll"
	pollsvc "pollparlor/internal/service/poll"
)

func main() {
	repo := pollrepo.NewMemoryRepo(seedPolls())
	service := pollsvc.New(repo)
	h := handler.NewPollHandler(service)
	r := router.New(h)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown: %v", err)
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
