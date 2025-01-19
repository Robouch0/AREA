//
// EPITECH PROJECT, 2025
// AREA
// File description:
// httpUtils_test
//

package http_utils

import (
	"net/http"
	"testing"
)

const (
	freeJSONAPIurl = "https://jsonplaceholder.typicode.com/todos/1"
)

func TestDefaultHeader(t *testing.T) {
	tok := "miaou"
	h1 := GetDefaultHTTPHeader(tok)
	if h1["Authorization"][0] != tok {
		t.Errorf("Token was not registered")
	}
}

func TestDefaultBotHeader(t *testing.T) {
	tok := "miaou"
	h1 := GetDefaultBotHTTPHeader(tok)
	if h1["Authorization"][0] != "Bot "+tok {
		t.Errorf("Token was not registered")
	}
}

func TestDefaultBearerHeader(t *testing.T) {
	tok := "miaou"
	h1 := GetDefaultBearerHTTPHeader(tok)
	if h1["Authorization"][0] != "Bearer "+tok {
		t.Errorf("Token was not registered")
	}
}

func TestSendHttpRequest(t *testing.T) {
	t.Run("Default JSON API URL", func(t *testing.T) {
		req, _ := http.NewRequest("GET", freeJSONAPIurl, nil)
		_, err := SendHttpRequest(req, 200)

		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
	t.Run("Wrong URL", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "uwu.com", nil)
		_, err := SendHttpRequest(req, 200)

		if err == nil {
			t.Errorf("Error was not catched: %v", err)
		}
	})
	t.Run("Wrong URL", func(t *testing.T) {
		req, _ := http.NewRequest("GET", freeJSONAPIurl, nil)
		_, err := SendHttpRequest(req, 204)

		if err == nil {
			t.Errorf("Error was not catched: %v", err)
		}
	})
}

func TestWriteHTTPResponse(t *testing.T) {
	t.Run("Empty response", func(t *testing.T) {
		WriteHTTPResponseErr(nil, 401, "Uwu")
	})
}
