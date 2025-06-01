package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DevVictor19/pic-pay-challenge/internal/domain/auth"
	"github.com/DevVictor19/pic-pay-challenge/internal/domain/user"
	"github.com/DevVictor19/pic-pay-challenge/internal/domain/wallet"
	"github.com/DevVictor19/pic-pay-challenge/internal/infra/db"
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

	database, err := db.Connect(
		cfg.DB.URL,
		cfg.DB.MaxOpenConns,
		cfg.DB.MaxIdleConns,
		cfg.DB.MaxIdleTime,
	)
	if err != nil {
		return err
	}
	defer database.Close()

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
	cfg, err := env.GetEnv()
	if err != nil {
		panic(err)
	}

	database, err := db.Get()
	if err != nil {
		panic(err)
	}

	userRepo := user.NewUserRepository(database, db.QueryDuration)
	userService := user.NewUserService(userRepo)

	walletRepo := wallet.NewWalletRepository(database, db.QueryDuration)
	walletService := wallet.NewWalletService(walletRepo)

	jwtService := auth.NewJWTService(cfg.JWT.Secret, cfg.JWT.Aud, cfg.JWT.Iss)
	bcryptService := auth.NewBcryptService()
	authService := auth.NewAuthService(userService, walletService, bcryptService, jwtService)
	authHandler := auth.NewAuthHandler(authService)

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

		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", utils.MakeHandler(authHandler.Signup))
			r.Post("/login", utils.MakeHandler(authHandler.Login))
		})

		// protected routes
		r.Group(func(r chi.Router) {
			r.Use(MakeJWTAuthMiddleware(jwtService, userService))

			r.Route("/transactions", func(r chi.Router) {
				r.Post("/", func(w http.ResponseWriter, r *http.Request) {
					utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "ok"})
				})
			})
		})
	})

	return r
}
