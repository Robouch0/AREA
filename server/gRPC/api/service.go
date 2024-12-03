//
// EPITECH PROJECT, 2024
// AREA
// File description:
// service
//

package api

type ClientService interface {
	SendAction(body []byte) (string, error)
	// TriggerReaction(action int, prevRes string) (int, error)
}
