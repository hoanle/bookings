package handlers

import (
	"net/http"

	"github.com/hoanle/bookings/pkg/config"
	"github.com/hoanle/bookings/pkg/models"
	"github.com/hoanle/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) About(rw http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	stringMap["remote_ip"] = repo.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
