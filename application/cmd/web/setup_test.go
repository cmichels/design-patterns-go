package main

import (
	"breeders/adapters"
	"breeders/configuration"
	"os"
	"testing"
)


var testApp application



func TestMain(m *testing.M){
  testBackend := &adapters.TestBackend{}
  testAdaptor := &adapters.RemoteService{Remote: testBackend}
  testApp = application{
    App: configuration.New(nil, testAdaptor),
  }


  os.Exit(m.Run())
}


