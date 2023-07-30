package main

import (
	"fmt"
	"reflect"
	"strings"

	genericrepository "github.com/rizkyrwijaya/all-about-golang/GoLearning/Generics/generic-interface/genericRepository"
	"github.com/rizkyrwijaya/all-about-golang/GoLearning/Generics/generic-interface/genericRepository/http"
	sourcea "github.com/rizkyrwijaya/all-about-golang/GoLearning/Generics/generic-interface/modelCollection/sourceA"
	httpimplementation "github.com/rizkyrwijaya/all-about-golang/GoLearning/Generics/generic-interface/normalRepository/httpImplementation"
)

// Repo Collections
var SourceARepo genericrepository.Requester[sourcea.TheModel]

func main() {
	// Using Normal Repository
	repo := httpimplementation.New("url", []byte(`{"data":"hellow"}`))
	_ = repo.GetResponse()
	res := repo.MockResponse()
	fmt.Println(res)

	fmt.Println(strings.Repeat("=", 80))

	// Using Generic Repository
	type Data struct {
		Test string `json:"testing"`
	}

	// As seen we can just add the data in the same file
	genericRepo := http.New[Data]("testUrl", []byte(`{"testing":"hello","data":"hellowww"}`))
	_ = genericRepo.GetResponse()
	res2 := genericRepo.MockResponse()
	fmt.Println(res2)
	fmt.Println(reflect.TypeOf(res2))

	// Or create a seperate package.
	SourceARepo = http.New[sourcea.TheModel]("testSourceA", []byte(`{"id":1234,"data":"insertDataHere"}`))
	_ = SourceARepo.GetResponse()
	res3 := SourceARepo.MockResponse()
	fmt.Println(res3)
	fmt.Println(reflect.TypeOf(res3))

}
