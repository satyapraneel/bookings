package middlewares

import (
	"fmt"
	"net/http"

	"github.com/satyapraneel/bookings/pkg/config"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("page")
		next.ServeHTTP(w, r)
		fmt.Println("page")
	})
}

func NoSurf(app *config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		csrfHandler := nosurf.New(next)
		csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: true,
			Secure:   app.InProduction,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
		})
		return csrfHandler
	}
}

func SessionLoad(app *config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return app.Session.LoadAndSave(next)
	}

}
