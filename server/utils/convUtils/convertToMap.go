//
// EPITECH PROJECT, 2025
// AREA
// File description:
// convertToMap
//

package conv_utils

import "encoding/json"

func ConvertToMap[T any](data *T) map[string]any {
	var m map[string]any

	b, err := json.Marshal(data)
	if err != nil {
		return map[string]any{}
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return map[string]any{}
	}
	return m
}
