package handlers

import (
	"net/http"

	"github.com/Alexis3386/bookings/pkg/config"
	"github.com/Alexis3386/bookings/pkg/models"
	"github.com/Alexis3386/bookings/pkg/render"
)

// TemplateData holds data send from handlers to template

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

// Home is the home page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.Template(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, test."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.Template(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is a Go function that handles making reservations.
//
// It takes a http.ResponseWriter and a http.Request as parameters and does not return anything.
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.Template(w, "make-reservation.page.html", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.Template(w, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	render.Template(w, "majors.page.html", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.Template(w, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.Template(w, "contact.page.html", &models.TemplateData{})
}
