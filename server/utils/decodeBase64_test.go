//
// EPITECH PROJECT, 2025
// AREA
// File description:
// decodeBase64_test
//

package utils

import (
	"encoding/json"
	"testing"
)

type person struct {
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	Address string `json:"address"`
}

type wrongstruct struct {
	Uwu string `json:"uwu"`
}

func TestDecodeBase64ToStruct(t *testing.T) {
	tests := []struct {
		name   string
		Person person `json:"person"`
	}{
		{
			name:   "Empty Person",
			Person: person{},
		},
		{
			name: "Real Person",
			Person: person{
				Name:    "Gragrou",
				Age:     17,
				Address: "miaou",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(&tt.Person)
			if err != nil {
				t.Errorf("Cannot transform to byte the data: %v", err)
				return
			}
			p2, err := DecodeBase64ToStruct[person](EncodeToBase64Bytes(string(b)))
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			if tt.Person != *p2 {
				t.Errorf("Cannot decode from base64")
			}
		})
	}

	t.Run("Failing incorrect bytes", func(t *testing.T) {
		_, err := DecodeBase64ToStruct[person]([]byte("Mamacita"))
		if err == nil {
			t.Errorf("Error is not catched")
			return
		}
	})
	b, err := json.Marshal(&tests[0].Person)
	if err != nil {
		t.Errorf("Cannot transform to byte the data: %v", err)
		return
	}
	t.Run("Failing incorrect bytes", func(t *testing.T) {
		_, err := DecodeBase64ToStruct[wrongstruct](EncodeToBase64Bytes(string(b)))
		if err != nil {
			t.Errorf("Error is not catched")
			return
		}
	})
}

func TestDecodeBase64ToString(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{
			name: "big str",
			str:  "LALALALALLALALALALALALLA",
		},
		{
			name: "Empty str",
			str:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p2, err := DecodeBase64ToString(EncodeToBase64Bytes(tt.str))
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			if tt.str != p2 {
				t.Errorf("Cannot decode from base64")
			}
		})
	}
}

func TestEncodeString(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{
			name: "Empty encode",
			str:  "",
		},
		{
			name: "Default encode",
			str:  "Str to encode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := EncodeToBase64(tt.str)
			dec, _ := DecodeBase64ToString([]byte(enc))
			if dec != tt.str {
				t.Errorf("Error on encode decode")
			}
		})
	}
}
