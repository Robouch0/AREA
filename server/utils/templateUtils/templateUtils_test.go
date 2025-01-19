//
// EPITECH PROJECT, 2025
// AREA
// File description:
// templateUtils_test
//

package template_utils

import "testing"

func TestCreateTemplate(t *testing.T) {
	t.Run("Default Template", func(t *testing.T) {
		_, err := CreateNewTemplate("MiaouGangster", "{{.Nya}} {{.Grrr}}")
		if err != nil {
			t.Errorf("Invalid template: %v", err)
		}
	})

	t.Run("Invalid Template", func(t *testing.T) {
		_, err := CreateNewTemplate("MiaouGangster", "{{Nya}} {{.Grrr}}")
		if err == nil {
			t.Errorf("Error was not catched")
		}
	})
}

func TestFormatIngredients(t *testing.T) {
	t.Run("Default Template", func(t *testing.T) {
		ingredients := map[string]any{
			"Dragon du 27": "{{.Count}} {{.MAAAA}}",
		}
		cache := map[string]any{
			"Count": "23",
			"MAAAA": 1223,
		}
		FormatIngredients(&ingredients, &cache)
		if ingredients["Dragon du 27"] != "23 1223" {
			t.Errorf("Invalid format for cache merging")
		}
	})

	t.Run("Invalid Template", func(t *testing.T) {
		ingredients := map[string]any{
			"Dragon du 27": "{{Count}} {{MAAAA}}",
		}
		cache := map[string]any{
			"Count": "23",
			"MAAAA": 1223,
		}
		FormatIngredients(&ingredients, &cache)
		if ingredients["Dragon du 27"] != "{{Count}} {{MAAAA}}" {
			t.Errorf("Invalid format for cache merging")
		}
	})
}
