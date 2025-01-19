//
// EPITECH PROJECT, 2025
// AREA
// File description:
// convertToMap_test
//

package conv_utils

import "testing"

type personTest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestConvertToMap(t *testing.T) {
	t.Run("Default conversion", func(t *testing.T) {
		m := ConvertToMap(&personTest{
			Name: "Graoul",
			Age:  17,
		})
		if _, ok := m["name"]; !ok {
			t.Errorf("Conversion did not work")
		}
		if _, ok := m["age"]; !ok {
			t.Errorf("Conversion did not work")
		}
	})
	t.Run("Empty conversion", func(t *testing.T) {
		m := ConvertToMap(&personTest{})
		if _, ok := m["name"]; !ok {
			t.Errorf("Conversion did not work")
		}
		if _, ok := m["age"]; !ok {
			t.Errorf("Conversion did not work")
		}
	})

	t.Run("Empty data", func(t *testing.T) {
		ConvertToMap[personTest](nil)
	})
}
