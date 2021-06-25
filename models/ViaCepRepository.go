package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Localidade  string `json:"localidade,omitempty"`
	Uf          string `json:"uf,omitempty"`
}

type Address struct {
	PostalCode   string `json:"postalCode"`
	Street       string `json:"street,omitempty"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
}

func ViaCepGetPostalCode(postalCode string) (Address, error) {
	log.Println("Repo: (ViaCep) Get - postalCode: ", postalCode)
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", postalCode)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	var address Address
	address.PostalCode = responseObject.Cep
	address.Street = responseObject.Logradouro
	address.Complement = responseObject.Complemento
	address.Neighborhood = responseObject.Localidade
	address.State = responseObject.Uf

	log.Printf("Repo: (Address) Get - body Out= %+v\n", address)
	return address, nil
}
