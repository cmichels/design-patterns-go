package adapters


import (
	"breeders/models"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
  GetBreedByName(breed string) (*models.CatBreed, error)
}

type RemoteService struct {
	Remote CatBreedsInterface
}

func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}
func (rs *RemoteService) GetBreedByName(breed string) (*models.CatBreed, error){
  return rs.Remote.GetBreedByName(breed)
}

type JSONBackend struct{}

func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")

	if err != nil {
		return nil, err
	}
 
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var breeds []*models.CatBreed
	err = json.Unmarshal(body, &breeds)

	if err != nil {
		return nil, err
	}

	return breeds, nil

}


func (rs *JSONBackend) GetBreedByName(b string) (*models.CatBreed, error){
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/json")

	if err != nil {
		return nil, err
	}
 
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var breed models.CatBreed
	err = json.Unmarshal(body, &breed)

	if err != nil {
		return nil, err
	}

	return &breed, nil
}




type XMLBackend struct{}

func (xb *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
  resp,err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")

  log.Println("xml")
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  body, err := io.ReadAll(resp.Body)

  if err != nil {
    return nil, err
  }


  type catBreeds struct {
    XMLName struct{} `xml:"cat-breeds"`
    Breeds []*models.CatBreed `xml:"cat-breed"`
  }

  var breeds catBreeds

  err = xml.Unmarshal(body, &breeds)

  return breeds.Breeds, nil

}

func (xb *XMLBackend) GetBreedByName(b string) (*models.CatBreed, error) {
  resp,err := http.Get("http://localhost:8081/api/cat-breeds/" + b + "/xml")

  log.Println("xml")
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  body, err := io.ReadAll(resp.Body)

  if err != nil {
    return nil, err
  }

  var breed models.CatBreed

  err = xml.Unmarshal(body, &breed)

  return &breed, nil

}


type TestBackend struct {}


func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error){
  return []*models.CatBreed{
    &models.CatBreed{ID:1, Breed: "tomcat", Details: "come d"},
  },nil
}

func (tb *TestBackend) GetBreedByName(b string) (*models.CatBreed, error){
  return &models.CatBreed{ID:1, Breed: b, Details: "come d"},nil
}
