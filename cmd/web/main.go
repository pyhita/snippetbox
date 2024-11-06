package main

import (
	"flag"
	"github.com/pyhita/snippetbox/cmd/web/handlers"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

func main() {

	//addr := flag.String("addr", ":4000", "HTTP listen address")
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "http service address")
	flag.StringVar(&cfg.staticDir, "static", "./static", "http service static dir")

	flag.Parse()

	//f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//// 记录 Info log 到文件中
	//infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	application := &handlers.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	// 创建自定义的http server，以便于可以自定义log 记录
	srv := &http.Server{
		Addr:     cfg.addr,
		Handler:  application.Routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s", cfg.addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
