//
// EPITECH PROJECT, 2025
// AREA
// File description:
// mergeMaps_test
//

package utils

import "testing"

func TestMergeMaps(t *testing.T) {
	t.Run("Merge map default test", func(t *testing.T) {
		dest := &map[string]string{
			"Miaou": "draadreaa",
			"Mew":   "drouuu",
			"Ippy":  "HHAHAHAH",
		}
		src := &map[string]string{
			"Drip check": "Yay",
		}
		MergeMaps[string](dest, src)
		if _, ok := (*dest)["Drip check"]; !ok {
			t.Errorf("Merge did not work")
		}
	})
}
