//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// getCryptoInfos
//

package cryptocompareserver

import (
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	cryptoAPIURL = "https://min-api.cryptocompare.com/data/price?fsym=%v&tsyms=%v"
)

func GetCryptoInfos(config *CryptoConfig) (*CryptoAPIResponseBody, error) {
	Url := fmt.Sprintf(cryptoAPIURL, config.CryptoCurrency, config.Currency)
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		log.Println("Request creation error: ", err)
		return nil, status.Errorf(codes.Internal, "Could not create the request: %v", err)
	}
	resp, err := http_utils.SendHttpRequest(req, 200)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error : reading body %s", err)
	}

	var m map[string]float64
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return &CryptoAPIResponseBody{Currency: m}, nil
}
