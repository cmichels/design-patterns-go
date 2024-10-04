package models

import "database/sql"


type Repository interface{
  AllDogBreeds() ([]*DogBreed,error)
  GetBreedByName(breed string) (*DogBreed, error)
  GetDogOfMonthById(id int) (*DogOfMonth, error)
}


type mysqlRepo struct {
  DB *sql.DB
}


func newMysqlRepo(conn *sql.DB) Repository{
  return &mysqlRepo{
    DB: conn,
  }
}


type testRepo struct {
  DB *sql.DB
}


func newTestRepo(conn *sql.DB) Repository{
  return &testRepo{
    DB: nil,
  }
}

