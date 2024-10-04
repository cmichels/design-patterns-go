package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))


  mux.Get("/dog-of-month", app.DogOfMonth)

	mux.Get("/test-patterns", app.TestPatterns)

	mux.Get("/api/dog-from-factory", app.CreateDog)
	mux.Get("/api/cat-from-factory", app.CreateCat)
	mux.Get("/api/dog-from-afactory", app.CreateADog)
	mux.Get("/api/cat-from-afactory", app.CreateACat)
	mux.Get("/api/dog-breeds", app.GetAllDogBreeds)
  mux.Get("/api/cat-breeds", app.GetAllCatBreeds)
	mux.Get("/api/dog-from-builder", app.CreateBuilderDog)
	mux.Get("/api/cat-from-builder", app.CreateBuilderCat)
	mux.Get("/api/animal-from-factory/{species}/{breed}", app.AnimalFromFactory)

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)
	return mux
}
