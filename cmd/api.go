package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application  struct {
	config config //dependencie
}

// run
// mount
func (app *application) mount() http.Handler{
	r:= chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})


	//http.ListenAndServe(":3333", r)
	return  r
}

func (app *application) run(h http.Handler) error{
	srv:= &http.Server{
		Addr: app.config.addr,
		Handler:h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	log.Printf("server started in add %s",app.config.addr)

	return srv.ListenAndServe()
}


type config struct {
	addr string //addres, eg port8080
	db dbConfig
}


type dbConfig struct {
	dsn string
}