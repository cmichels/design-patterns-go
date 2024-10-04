package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)



func Test_getAllDogBreeds(t *testing.T) {

  req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(testApp.GetAllDogBreeds)
  handler.ServeHTTP(rr, req)

  if rr.Code != http.StatusOK{
    t.Errorf("inccorect status %d expect %d", rr.Code, http.StatusOK)
  } 

}


func Test_getAllCatBreeds(t *testing.T) {

  req, _ := http.NewRequest("GET", "/api/cat-breeds", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(testApp.GetAllCatBreeds)
  handler.ServeHTTP(rr, req)

  if rr.Code != http.StatusOK{
    t.Errorf("inccorect status %d expect %d", rr.Code, http.StatusOK)
  } 

}
