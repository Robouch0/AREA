//
// EPITECH PROJECT, 2025
// AREA
// File description:
// formatIngredients
//

package template_utils

import (
	"reflect"
	"strings"
)

func formatStringIngredient(k, v string, src *map[string]any) string {
	t, err := CreateNewTemplate(k, v)
	if err != nil {
		return v // If format is invalid maybe it is not a format ^^
	}
	builder := &strings.Builder{}
	t.Execute(builder, *src)
	return builder.String()
}

func FormatIngredients(dest, src *map[string]any) {
	for k, v := range *dest {
		if reflect.TypeOf(v).Kind() == reflect.String {
			(*dest)[k] = formatStringIngredient(k, v.(string), src)
		}
	}
}
