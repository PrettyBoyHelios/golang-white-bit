package models

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Whitebit struct {
	PublicKey string
	SecretKey string
	BaseURL   string
}

func (w *Whitebit) MarketInfo() (market []Market, err error) {
	endpoint := "api/v2/public/markets"
	resp, err := w.sendRequestUnmarshal(w.BaseURL+endpoint, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Result, &market)
	if err != nil {
		return
	}
	return
}

func (w *Whitebit) MarketActivity() (marketActivity []MarketActivity, err error) {
	endpoint := "api/v2/public/markets"
	resp, err := w.sendRequestUnmarshal(w.BaseURL+endpoint, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Result, &marketActivity)
	if err != nil {
		return
	}
	return
}

func (w *Whitebit) sendRequest(requestURL string, data map[string]string) (responseBody []byte, err error) {
	//If the nonce is similar to or lower than the previous request number, you will receive the 'too many requests' error message
	nonce := int(time.Now().Unix()) //nonce is a number that is always higher than the previous request number

	data["request"] = requestURL
	data["nonce"] = strconv.Itoa(nonce)

	requestBody, err := json.Marshal(data)
	if err != nil {
		return
	}

	//preparing request URL
	completeURL := w.BaseURL + requestURL

	//calculating payload
	payload := base64.StdEncoding.EncodeToString(requestBody)

	//calculating signature using sha512
	h := hmac.New(sha512.New, []byte(w.SecretKey))
	h.Write([]byte(payload))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	client := http.Client{}

	request, err := http.NewRequest("POST", completeURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	//setting neccessarry headers
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("X-TXC-APIKEY", w.PublicKey)
	request.Header.Set("X-TXC-PAYLOAD", payload)
	request.Header.Set("X-TXC-SIGNATURE", signature)

	//sending request
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	//reciving data
	responseBody, err = ioutil.ReadAll(response.Body)

	return
}

func (w *Whitebit) sendRequestUnmarshal(requestURL string, data map[string]string) (res BaseResponse, err error) {
	//If the nonce is similar to or lower than the previous request number, you will receive the 'too many requests' error message
	nonce := int(time.Now().Unix()) //nonce is a number that is always higher than the previous request number

	data["request"] = requestURL
	data["nonce"] = strconv.Itoa(nonce)

	requestBody, err := json.Marshal(data)
	if err != nil {
		return
	}

	//preparing request URL
	completeURL := w.BaseURL + requestURL

	//calculating payload
	payload := base64.StdEncoding.EncodeToString(requestBody)

	//calculating signature using sha512
	h := hmac.New(sha512.New, []byte(w.SecretKey))
	h.Write([]byte(payload))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	client := http.Client{}

	request, err := http.NewRequest("POST", completeURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	//setting neccessarry headers
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("X-TXC-APIKEY", w.PublicKey)
	request.Header.Set("X-TXC-PAYLOAD", payload)
	request.Header.Set("X-TXC-SIGNATURE", signature)

	//sending request
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	//reciving data
	responseBody, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(responseBody, &res)

	if !res.Success {
		return res, errors.New("unsuccessful request")
	}

	return
}
