//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// cryptoTypes
//

package cryptocompareserver

type CryptoConfig struct {
	CryptoCurrency string  `json:"crypto_currency"`
	Currency       string  `json:"currency"`
}

type CryptoAPIResponseBody struct {
	Currency       map[string]float64 `json:"currency,omitempty"`
}
