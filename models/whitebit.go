package models

import (
	"fmt"
	"log"
)

type publicAPI struct {

}

type privateAPI struct {
	hasPrivate bool
	apiKey string
}

func (p *privateAPI) GetAccountBalance() {
	if p.hasPrivate {

	} else {
		log.fAT
	}
	fmt.Print("retrieving balance")
}

type Whitebit struct {
	publicAPI
	privateAPI
}

func NewWhitebit(ApiKey string) *Whitebit {
	w := new(Whitebit)
	w.privateAPI.apiKey = ApiKey
	if ApiKey == ""{
		w.hasPrivate = false
		log.Println("Private API is disabled")
	} else {
		w.hasPrivate = true
	}
	return w
}