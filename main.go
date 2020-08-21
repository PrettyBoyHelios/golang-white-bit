package main

import "github.com/PrettyBoyHelios/golang-white-bit/models"

func main() {

	whitebitApi := models.NewWhitebit("")
	whitebitApi.GetAccountBalance()
}