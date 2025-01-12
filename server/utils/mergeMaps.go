//
// EPITECH PROJECT, 2025
// AREA
// File description:
// mergeMaps
//

package utils

func MergeMaps[T any](dest, src *map[string]T) {
	for k, v := range *src {
		(*dest)[k] = v
	}
}
