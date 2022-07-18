package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/satyapraneel/bookings/pkg/config"
	"github.com/satyapraneel/bookings/pkg/render"
	"github.com/satyapraneel/bookings/routes"

	"github.com/alexedwards/scs/v2"
)

const PORT_NUMBER = ":8000"

func main() {
	//

	app := setupApp()
	mux := routes.Routes(app)

	fmt.Println("server started at " + PORT_NUMBER)
	err := http.ListenAndServe(PORT_NUMBER, mux)
	panic(err)
}

func setupApp() *config.AppConfig {
	var app config.AppConfig
	render.NewTemplate(&app)
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	return &app
}
