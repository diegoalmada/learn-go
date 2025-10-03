package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AwesomeZipCode struct {
	Cep         string `json:"cep"`
	AddressType string `json:"address_type"`
	AddressName string `json:"address_name"`
	Address     string `json:"address"`
	State       string `json:"state"`
	District    string `json:"district"`
	Lat         string `json:"lat"`
	Lng         string `json:"lng"`
	City        string `json:"city"`
	CityIbge    string `json:"city_ibge"`
	Ddd         string `json:"ddd"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao obter url %s: %v\n", url, err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao ler resposta de %s: %v\n", url, err)
		}
		var data AwesomeZipCode
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao desserializar resposta de %s: %v\n", url, err)
		}
		file, err := os.Create("zipcode.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao criar arquivo zipcode.txt: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s\n", data.Cep, data.City))
	}
}
