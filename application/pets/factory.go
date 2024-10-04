package pets

import "breeders/models"


// Factory to create base pet (dog/cat) with default values
func NewPet(species string) *models.Pet{
  pet := models.Pet{
    Species: species,
    Breed: "",
    MinWeight: 0,
    MaxWeight: 0,
    Desc: "no desc",
    LifeSpan: 0,
  }

  return &pet
}
