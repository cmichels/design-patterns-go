package main

import (
	"breeders/adapters"
	"breeders/configuration"
	"breeders/streamer"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	port       = ":4000"
	timeoutDur = 30 * time.Second
)

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
  videoQueue chan streamer.VideoProcessingJob
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {

  const numWorkers = 4

  videoQueue := make(chan streamer.VideoProcessingJob, numWorkers)
  defer close(videoQueue)

	app := application{
		templateMap: make(map[string]*template.Template),
    videoQueue: videoQueue,
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsb", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

  // jsonBackend := &adapters.JSONBackend{}
  // jsonAdapter := &adapters.RemoteService{Remote: jsonBackend}

  xmlBackend := &adapters.XMLBackend{}
  xmlAdapter := &adapters.RemoteService{Remote: xmlBackend}

	if db, err := initDb(app.config.dsn); err != nil {
		log.Panic(err)
	} else {
		app.App =  configuration.New(db, xmlAdapter)
	}


  wp := streamer.New(videoQueue, numWorkers)
  wp.Run()

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       timeoutDur,
		ReadTimeout:       timeoutDur,
		ReadHeaderTimeout: timeoutDur,
		WriteTimeout:      timeoutDur,
	}

	fmt.Println("starting on part", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Panicf("failed to start server. cause by: %s", err)
	}
}
