//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroGenerics
//

package mirotypes

type MiroGenericBody[T any] struct {
	Data T `json:"data,omitempty"`
}
