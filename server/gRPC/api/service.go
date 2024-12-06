//
// EPITECH PROJECT, 2024
// AREA
// File description:
// service
//

package api

type ClientService interface {
	SendAction(body map[string]any) (string, error)
	// TriggerReaction(action int, prevRes string) (int, error)
}
