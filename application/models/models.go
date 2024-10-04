package models

import (
	"database/sql"
	"time"
)

var repo Repository

type Models struct {
	DogBreed DogBreed
  Dog Dog
}

func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = newMysqlRepo(conn)
	} else {
		repo = newTestRepo(nil)
	}

	return &Models{
		DogBreed: DogBreed{},
	}
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return repo.AllDogBreeds()
}

func (d *DogBreed) GetBreedByName(breed string) (*DogBreed, error) {
	return repo.GetBreedByName(breed)
}

func (d *Dog) GetDogOfMonthById(id int) (*DogOfMonth , error) {
	return repo.GetDogOfMonthById(id)
}

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	LifeSpan         int    `json:"life_span"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origing"`
}
type CatBreed struct {
	ID               int    `json:"id" xml:"id"`
	Breed            string `json:"breed" xml:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs" xml:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs" xml:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight" xml:"average_weight"`
	LifeSpan         int    `json:"life_span" xml:"life_span"`
	Details          string `json:"details" xml:"details"`
	AlternateNames   string `json:"alternate_names" xml:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin" xml:"geographic_origin"`
}

type Cat struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	Dob              time.Time `json:"dob"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Desc             string    `json:"desc"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}
type Dog struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	Dob              time.Time `json:"dob"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Desc             string    `json:"desc"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}
type Breeder struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Address   string      `json:"address"`
	City      string      `json:"city"`
	State     string      `json:"state"`
	Country   string      `json:"country"`
	Zip       string      `json:"zip"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Active    int         `json:"active"`
	DogBreeds []*DogBreed `json:"dog_breeds"`
	CatBreeds []*CatBreed `json:"cat_breeds"`
}
type Pet struct {
	Species   string `json:"species"`
	Breed     string `json:"breed"`
	MinWeight int    `json:"min_weight"`
	MaxWeight int    `json:"max_weight"`
	Desc      string `json:"desc"`
	LifeSpan  int    `json:"life_span"`
}


type DogOfMonth struct {
  ID int
  Dog *Dog
  Video string
  Image string
}
