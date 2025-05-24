package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/env"
	"github.com/DevVictor19/pic-pay-challenge/internal/infra/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() error {
	cfg, err := env.LoadEnv()
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerPort),
		Handler:      mount(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "ok"})
		})
	})

	return r
}
