package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/justinas/alice"
)

// 统一注册路由
func (a *Application) Routes() http.Handler {

	// match url and handler
	router := httprouter.New()

	// custom not found handler
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static/", fileServer))

	router.HandlerFunc(http.MethodGet, "/", a.Home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", a.SnippetView)
	router.HandlerFunc(http.MethodPost, "/snippet/create", a.SnippetCreate)

	// Create a middleware chain containing our 'standard' middlewares
	standard := alice.New(a.recoverPanic, a.logRequest, a.secureHeader)

	return standard.Then(router)
}
