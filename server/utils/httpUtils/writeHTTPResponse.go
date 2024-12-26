//
// EPITECH PROJECT, 2024
// AREA
// File description:
// writeHTTPResponseErr
//

package http_utils

import "net/http"

func WriteHTTPResponseErr(w *http.ResponseWriter, code int, err string) {
	if w == nil {
		return
	}
	(*w).WriteHeader(code)
	(*w).Write([]byte(err))
}
