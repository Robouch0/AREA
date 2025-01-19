//
// EPITECH PROJECT, 2025
// AREA
// File description:
// getEnvParam_test
//

package utils

import "testing"

func TestGetEnvParameter(t *testing.T) {
	tests := []struct {
		name string
		key  string
	}{
		{
			name: "Secret key",
			key:  "SECRET_KEY",
		},
		{
			name: "Secret key",
			key:  "DATABASE_URL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetEnvParameter(tt.key)
		})
	}
}
