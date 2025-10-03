package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ZipCode struct {
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", findZipCodeInvoke)
	mux.Handle("/blog", blog{title: "Blog do GoLand"})
	mux.Handle("/golang", blog{title: "Blog do Golang"})
	http.ListenAndServe(":8082", mux)
}

func findZipCodeInvoke(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	param := r.URL.Query().Get("zipcode")
	if param == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	zipcode, err := findZipCode(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(zipcode)
}

func findZipCode(zipcode string) (*ZipCode, error) {
	req, err := http.Get("https://cep.awesomeapi.com.br/json/" + zipcode)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var z ZipCode
	err = json.Unmarshal(body, &z)
	if err != nil {
		return nil, err
	}
	return &z, nil
}

type blog struct {
	title string
}

// Vantagem de usar esse approach.
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
