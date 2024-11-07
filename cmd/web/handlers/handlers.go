package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pyhita/snippetbox/internal/models"
)

func (a *Application) Home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/".
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

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
	// read snippet id from query string
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	fmt.Fprintf(w, "Display specific snippet %v", snippet)
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
