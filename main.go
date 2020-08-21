package main

import (
	"fmt"
	"github.com/PrettyBoyHelios/golang-white-bit/models"
	"log"
	"os"
)

func main() {

	whitebit := models.Whitebit{
		PublicKey: os.Getenv("WB_PUBLIC_KEY"),
		SecretKey: os.Getenv("WB_PRIVATE_KEY"),
		BaseURL:   "https://whitebit.com",
	}

	resultData, err := whitebit.MarketInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultData)

	resultWithdraw, err := whitebit.Withdraw(models.WithdrawParams{
		Ticker:   "POLIS",
		Amount:   "4",
		Address:  "PjhsadasyuvbaJButHJvt",
		Memo:     "test",
		UniqueID: "728y73432",
		Request:  "",
		Nonce:    "83838",
	})
	fmt.Println(resultWithdraw)
}
