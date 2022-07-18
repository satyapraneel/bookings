package routes

import (
	"net/http"

	"github.com/satyapraneel/bookings/middlewares"
	"github.com/satyapraneel/bookings/pkg/config"
	"github.com/satyapraneel/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(app *config.AppConfig) http.Handler {

	repo := handlers.NewHandler(app)
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer, middleware.Logger)
	mux.Use(middlewares.NoSurf(app), middlewares.SessionLoad(app))
	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)
	return mux
}
