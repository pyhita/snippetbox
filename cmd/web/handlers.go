package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/pyhita/snippetbox/internal/models"
)

func (a *Application) Home(w http.ResponseWriter, r *http.Request) {
	snippets, err := a.Snippets.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "<h1>Snippet %s</h1>", snippet.Title)
	}
}

func (a *Application) SnippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		a.notFound(w)
		return
	}

	snippet, err := a.Snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			a.notFound(w)
		} else {
			a.serverError(w, err)
		}
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/view.tmpl",
		"./ui/html/partials/nav.tmpl",
	}
	// parse the template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, err)
		return
	}

	// 将所有需要传递给模板的数据封装到 templateData中
	data := &templateData{
		Snippet: snippet,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		a.serverError(w, err)
	}
}

func (a *Application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	// create snippet only supports post method
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		//w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write 默认设置status code 200，所以必须在之前WriteHeader
		a.clientError(w, http.StatusMethodNotAllowed)
		//http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := a.Snippets.Insert(title, content, expires)
	if err != nil {
		a.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%v", id), http.StatusSeeOther)
}
