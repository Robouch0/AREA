//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana Board service
//

package asana_generics

type AsanaBaseBody[T any] struct {
	Data T `json:"data,omitempty"`
}
