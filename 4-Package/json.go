package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{1, 100}
	_, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		fmt.Println(err)
	}

	//deserializar meu json
	pureJson := []byte(`{"numero": 2, "saldo": 80}`)

	var contaX Conta
	err = json.Unmarshal(pureJson, &contaX)
	if err != nil {
		panic(err)
	}
	fmt.Println(contaX.Saldo, contaX.Numero)
}
