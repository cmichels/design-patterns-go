package models

import (
	"context"
	"log"
	"time"
)

func (m *mysqlRepo) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
  cast(((weight_low_lbs + weight_high_lbs)/2) as unsigned) as average_weight,
  lifespan, coalesce(details,''),
  coalesce(alternate_names,''), coalesce(geographic_origin,'')
  from dog_breeds order by breed`

	var breeds []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var d DogBreed

		err := rows.Scan(&d.ID,
			&d.Breed, &d.WeightLowLbs, &d.WeightHighLbs, &d.AverageWeight,
			&d.LifeSpan, &d.Details, &d.AlternateNames, &d.GeographicOrigin)

		if err != nil {
			return nil, err
		}

		breeds = append(breeds, &d)
	}

	return breeds, nil
}

func (m *mysqlRepo) GetBreedByName(b string) (*DogBreed, error) {
  log.Println("get breed by name")
  log.Println(b)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
  cast(((weight_low_lbs + weight_high_lbs)/2) as unsigned) as average_weight,
  lifespan, coalesce(details,''),
  coalesce(alternate_names,''), coalesce(geographic_origin,'')
  from dog_breeds where breed = ?`


	row := m.DB.QueryRowContext(ctx, query, b)

	var d DogBreed

	err := row.Scan(&d.ID,
		&d.Breed, &d.WeightLowLbs, &d.WeightHighLbs, &d.AverageWeight,
		&d.LifeSpan, &d.Details, &d.AlternateNames, &d.GeographicOrigin)

	if err != nil {
		return nil, err
	}

	return &d, nil
}


func (m *mysqlRepo) GetDogOfMonthById(id int) (*DogOfMonth, error){

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, video, image from dog_of_month where id = ?`

	var d DogOfMonth

	row := m.DB.QueryRowContext(ctx, query, id)


	err := row.Scan(&d.ID,&d.Video,&d.Image)

	if err != nil {
		return nil, err
	}

	return &d, nil

}
