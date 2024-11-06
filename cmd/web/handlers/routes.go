package handlers

import "net/http"

// 统一注册路由

func (a *Application) Routes() *http.ServeMux {

	// match url and handler
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	// register the file server
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", a.Home)
	mux.HandleFunc("/snippet/view", a.SnippetView)
	mux.HandleFunc("/snippet/create", a.SnippetCreate)

	return mux
}
