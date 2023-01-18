package handlers

import (
	"log"
	"net/http"

	"github.com/yesilyurtburak/go-web-basics-3/models"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/config"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/forms"
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
	render.RenderTemplate(w, r, "home.page.gotmpl", &models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "about.page.gotmpl", &models.PageData{StrMap: strMap})
	// created a strMap and send some information to about.page.gotmpl template via models.PageData
}

func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "login.page.gotmpl", &models.PageData{StrMap: strMap})
}

func (m *Repository) PageHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "page.page.gotmpl", &models.PageData{StrMap: strMap})
}

func (m *Repository) MakePostHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "makepost.page.gotmpl", &models.PageData{
		Form: forms.NewForm(nil),
	})
}

func (m *Repository) PostMakePostHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm() // sets the r.Form and r.PostForm
	if err != nil {
		log.Fatal(err)
		return
	}

	// get the form data from template and create a new struct `article`
	article := models.Article{
		BlogTitle:   r.Form.Get("blog_title"),
		BlogArticle: r.Form.Get("blog_article"),
	}

	form := forms.NewForm(r.PostForm) // creates a new post form
	form.HasValue("blog_title", r)    // checks the field and add an error if it is empty
	form.HasValue("blog_article", r)  // checks the field and add an error if it is empty

	// if form has errors rerender the page and show the errors on the fields.
	if !form.IsValid() {
		data := make(map[string]interface{})
		data["article"] = article
		render.RenderTemplate(w, r, "makepost.page.gotmpl", &models.PageData{Form: form, DataMap: data})
		return
	}

	// // Access the form data
	// blog_title := r.Form.Get("blog_title")
	// blog_article := r.Form.Get("blog_article")

	// // Write the data to the response writer.
	// w.Write([]byte(blog_title))
	// w.Write([]byte(blog_article))
}
