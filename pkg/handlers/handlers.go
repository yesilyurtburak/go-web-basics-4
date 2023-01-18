package handlers

import (
	"net/http"

	"github.com/yesilyurtburak/go-web-basics-3/models"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/config"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/render"
)

// Type definition for Repository pattern
type Repository struct {
	App *config.AppConfig
}

// Variable declaration for Repository pattern
var Repo *Repository

// Function definition for creating a new Repository
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// Function definition to handle routing with Repository pattern
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gotmpl", &models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "about.page.gotmpl", &models.PageData{StrMap: strMap})
	// created a strMap and send some information to about.page.gotmpl template via models.PageData
}

func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "login.page.gotmpl", &models.PageData{StrMap: strMap})
}

func (m *Repository) MakePostHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "makepost.page.gotmpl", &models.PageData{StrMap: strMap})
}

func (m *Repository) PageHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, "page.page.gotmpl", &models.PageData{StrMap: strMap})
}
