package handlers

import "github.com/satyapraneel/bookings/pkg/config"

type Repository struct {
	App *config.AppConfig
}

func NewHandler(App *config.AppConfig) *Repository {
	return &Repository{
		App: App,
	}
}
