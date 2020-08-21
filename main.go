package main

import (
	"encoding/json"
	"log"
)

func main() {

	provider := Whitebit{
		PublicKey: "",
		SecretKey: "",
		BaseURL:   "https://whitebit.com",
	}

	//put here request path. For obtaining trading balance use: /api/v4/trade-account/balance
	request := "/api/v4/trade-account/balance"

	//put here data to send
	data := map[string]string{
		"ticker": "BTC", //for example for obtaining trading balance for BTC currency
	}

	resultData, err := provider.SendRequest(request, data)
	if err != nil {
		log.Fatal(err)
	}

	var result interface{}
	if err := json.Unmarshal(resultData, &result); err != nil {
		log.Fatal(err)
	}

	//printing response body to default output
	log.Println(result)
}
