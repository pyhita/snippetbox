package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// 统一注册路由
func (a *Application) Routes() http.Handler {

	// match url and handler
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	// register the file server
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", a.Home)
	mux.HandleFunc("/snippet/view", a.SnippetView)
	mux.HandleFunc("/snippet/create", a.SnippetCreate)

	// Create a middleware chain containing our 'standard' middlewares
	standard := alice.New(a.recoverPanic, a.logRequest, a.secureHeader)

	return standard.Then(mux)
}
