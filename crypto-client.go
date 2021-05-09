package client

import (
	"encoding/json"
	"log"
	"net/http"

	"https://github.com/grahamuk2018/go-crypto-price/model"
)

func FetchCrypto(fiat, crypto string) (string, error) {

	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + crypto + "&convert=" + fiat

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("Ooops an error occured, please try again")
	}
	defer resp.Body.Close()

	var cResp model.Cryptoresponse

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	return cResp.TextOutput(), nil
}
