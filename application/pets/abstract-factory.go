package pets

import (
	"breeders/configuration"
	"breeders/models"
	"errors"
	"fmt"
	"log"
)

type Animal interface {
	Show() string
}

type PetFactory interface {
	newPet() Animal
	newPetWithBreed(breed string) Animal
}

type DogFromFact struct {
	Pet *models.Dog
}

func (dff *DogFromFact) Show() string {
	return fmt.Sprintf("dog from fact: %v+", dff.Pet.Breed)
}

type CatFromFact struct {
	Pet *models.Cat
}

func (cff *CatFromFact) Show() string {
	return fmt.Sprintf("dog from facts %v+", cff.Pet.Breed)
}

type DogFactory struct{}

func (df *DogFactory) newPet() Animal {
	return &DogFromFact{
		Pet: &models.Dog{},
	}
}

func (df *DogFactory) newPetWithBreed(b string) Animal {
	app := configuration.GetInstance()
	breed, err := app.Models.DogBreed.GetBreedByName(b)

	if err != nil {
		log.Printf("error get dog %s", err)
		return nil
	}
	return &DogFromFact{
		Pet: &models.Dog{Breed: *breed},
	}
}

type CatFactory struct{}

func (df *CatFactory) newPet() Animal {
	return &CatFromFact{
		Pet: &models.Cat{},
	}
}
func (df *CatFactory) newPetWithBreed(b string) Animal {
	app := configuration.GetInstance()
	breed, err := app.CatService.GetBreedByName(b)

	if err != nil {
		return nil
	}

	return &CatFromFact{
		Pet: &models.Cat{Breed: *breed},
	}
}

func NewPetFrom(species string) (Animal, error) {
	switch species {
	case "dog":
		var f DogFactory
		return f.newPet(), nil

	case "cat":
		var f CatFactory
		return f.newPet(), nil
	default:
		return nil, errors.New("invalid species supplied")
	}

}

func NewPetWithBreed(species, breed string) (Animal, error) {
	switch species {
	case "dog":
		var dogFact DogFactory
		dog := dogFact.newPetWithBreed(breed)
		return dog, nil
	case "cat":
		var catFact CatFactory
		cat := catFact.newPetWithBreed(breed)
		return cat, nil
	default:
		return nil, errors.New("no pet available")
	}
}
