// Package bootstrap provides bootstrap for the application
package bootstrap

import (
	"context"
	"net/http"
	"time"

	"pollparlor/internal/config"
	"pollparlor/internal/http/handler"
	"pollparlor/internal/http/router"
	mongox "pollparlor/internal/infra/mongo"
	pollrepo "pollparlor/internal/repository/poll"
	userrepo "pollparlor/internal/repository/user"
	pollsvc "pollparlor/internal/service/poll"
	usersvc "pollparlor/internal/service/user"
)

// App represents the application
type App struct {
	Cleanup func(context.Context) error
	Server  *http.Server
}

// New creates a new App
func New(cfg *config.Config) (*App, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongox.NewClient(ctx, cfg.Mongo.URI)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := mongox.NewDatabase(client, cfg.Mongo.DB)

	// if err := mongox.EnsurePollIndexes(db.Collection("polls"), 20*time.Second); err != nil {
	// 	_ = client.Disconnect(context.Background())
	// 	return nil, err
	// }

	pollRepo := pollrepo.NewMongoRepo(db, 2*time.Second)
	userRepo := userrepo.NewMongoRepo(db, 2*time.Second)

	pollSvc := pollsvc.New(pollRepo)
	userSvc := usersvc.New(userRepo)

	pollHand := handler.NewPollHandler(pollSvc)
	userHand := handler.NewUserHandler(userSvc)

	r := router.New()
	apiV1 := r.Group("/api").Group("/v1")
	router.CreateHealthRoutes(apiV1)
	router.CreatePollRoutes(apiV1, pollHand)
	router.CreateUserRoutes(apiV1, userHand)

	srv := &http.Server{
		Addr:         cfg.App.Addr + ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &App{
		Cleanup: func(ctx context.Context) error {
			return client.Disconnect(ctx)
		},
		Server: srv,
	}, nil
}
