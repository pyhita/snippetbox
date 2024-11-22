package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/pyhita/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql" // New import
)

type config struct {
	addr      string
	staticDir string
	dsn       string
}

func main() {

	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "http service address")
	flag.StringVar(&cfg.staticDir, "static", "./static", "http service static dir")
	flag.StringVar(&cfg.dsn, "dsn", "root:root@tcp(127.0.0.1:3308)/snippetbox?parseTime=true", "database connect string")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create mysql connection pool
	db, err := openDB(cfg.dsn)
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	application := &Application{
		InfoLog:  infoLog,
		ErrorLog: errLog,
		Snippets: &models.SnippetModel{DB: db},
	}

	// 创建自定义的http server，以便于可以自定义log 记录
	srv := &http.Server{
		Addr:    cfg.addr,
		Handler: application.Routes(),
	}

	infoLog.Printf("Starting server on %s", cfg.addr)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
