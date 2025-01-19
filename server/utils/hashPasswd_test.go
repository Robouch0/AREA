//
// EPITECH PROJECT, 2025
// AREA
// File description:
// hashPasswd_test
//

package utils

import "testing"

func TestHashPassword(t *testing.T) {
	t.Run("Empty Password", func(t *testing.T) {
		_, err := HashPassword("")
		if err != nil {
			t.Errorf(err.Error())
		}
	})
	t.Run("Default Password", func(t *testing.T) {
		_, err := HashPassword("password")
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("Check default password", func(t *testing.T) {
		hash, err := HashPassword("password")
		if err != nil {
			t.Errorf(err.Error())
		}
		if CheckPasswordHash("password", hash) == false {
			t.Errorf("Not the same hash")
		}
	})
	t.Run("Empty pass", func(t *testing.T) {
		hash, err := HashPassword("")
		if err != nil {
			t.Errorf(err.Error())
		}
		if CheckPasswordHash("", hash) == false {
			t.Errorf("Not the same hash")
		}
	})
}
