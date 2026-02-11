package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	// converter STRUCT EM JSON
	c := cachorro{"Rex", "Damalta", 3}
	fmt.Println(c)

	cachorroJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroJSON)

	cachorroConvertido := bytes.NewBuffer(cachorroJSON)
	fmt.Println(cachorroConvertido)

}
