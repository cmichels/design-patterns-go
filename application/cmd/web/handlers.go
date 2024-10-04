package main

import (
	"breeders/models"
	"breeders/pets"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

var tb toolbox.Tools

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {

	page := chi.URLParam(r, "page")

	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDog(w http.ResponseWriter, r *http.Request) {
	_ = tb.WriteJSON(w, http.StatusOK, pets.NewPet("dog"), nil)
}

func (app *application) CreateCat(w http.ResponseWriter, r *http.Request) {
	_ = tb.WriteJSON(w, http.StatusOK, pets.NewPet("cat"), nil)
}

func (app *application) CreateADog(w http.ResponseWriter, r *http.Request) {

	dog, err := pets.NewPetFrom("dog")

	if err != nil {
		_ = tb.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	tb.WriteJSON(w, http.StatusOK, dog, nil)

}

func (app *application) CreateACat(w http.ResponseWriter, r *http.Request) {

	dog, err := pets.NewPetFrom("cat")

	if err != nil {
		_ = tb.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	tb.WriteJSON(w, http.StatusOK, dog, nil)

}

func (app *application) GetAllDogBreeds(w http.ResponseWriter, r *http.Request) {

	dogBreeds, err := app.App.Models.DogBreed.All()

	if err != nil {
		_ = tb.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = tb.WriteJSON(w, http.StatusOK, dogBreeds)

}

func (app *application) CreateBuilderDog(w http.ResponseWriter, r *http.Request) {

	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetAge(15).
		SetBreed("mixed").
		SetColor("black and white").
		SetWeight(100).
		SetLifeSpan(200).
		SetMaxWeight(300).
		SetMinWeight(100).
		SetDescription("mixed breed doggo").
		SetAgeEstimated(true).
		SetGeographicOrigin("japan").
		Build()

	if err != nil {
		_ = tb.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = tb.WriteJSON(w, http.StatusOK, p)

}

func (app *application) CreateBuilderCat(w http.ResponseWriter, r *http.Request) {

	p, err := pets.NewPetBuilder().
		SetSpecies("Cat").
		SetAge(15).
		SetBreed("mixed").
		SetColor("orange").
		SetWeight(100).
		SetLifeSpan(200).
		SetMaxWeight(300).
		SetMinWeight(100).
		SetDescription("mixed breed doggo").
		SetAgeEstimated(true).
		SetGeographicOrigin("japan").
		Build()

	if err != nil {
		_ = tb.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = tb.WriteJSON(w, http.StatusOK, p)

}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {

	catBreeds, err := app.App.CatService.GetAllBreeds()

	if err != nil {
		_ = tb.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = tb.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromFactory(w http.ResponseWriter, r *http.Request) {

	species := chi.URLParam(r, "species")
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)

	pet, err := pets.NewPetWithBreed(species, breed)

	if err != nil {
		tb.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = tb.WriteJSON(w, http.StatusOK, pet)

}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {

	breed, err := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")

  log.Printf("breed %v+", breed)
  if err != nil {
    log.Printf("breed error: %s", err)
    return
  }

	dom, err := app.App.Models.Dog.GetDogOfMonthById(1)

  if err != nil {
    log.Printf("dom error: %s", err)
    return
  }
  layout := "2006-01-02"
  dob,_ := time.Parse(layout, "2023-11-01")
  dog := models.DogOfMonth{
    Dog: &models.Dog{
      ID: 1,
      Name: "same",
      BreedID: breed.ID,
      Color: "black",
      Dob: dob,
      SpayedOrNeutered: 0,
      Desc: "good boy",
      Weight: 20,
      Breed: *breed,
    },
    Video: dom.Video,
    Image: dom.Image,
  }

  data := make(map[string]any)
  data["dog"] = dog

  app.render(w, "dog-of-month.page.gohtml", &templateData{Data: data})

}
