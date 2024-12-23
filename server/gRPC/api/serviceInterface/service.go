//
// EPITECH PROJECT, 2024
// AREA
// File description:
// service
//

package serviceinterface

import "area/models"

type ActionResponseStatus struct {
	Description string `json:"description"`
	ActionID    int    `json:"action_id"`
}

type ReactionResponseStatus struct {
	Description string `json:"description"`
	ReactionID  int    `json:"reaction_id"`
}

type MicroserviceLauncher = map[string]func(ingredients map[string]any, prevOutput []byte, userID int) (*ReactionResponseStatus, error)
type ActionLauncher = map[string]func(scenario models.AreaScenario, actionId, userID int) (*ActionResponseStatus, error)

// Map with the ingredient name mapped with his possible value type (named as a string)
//
// The value types possible are "string", "int", "time", "bool"
//
// Examples:
//   - Key: "Hour" | Value: "int"
//   - Key: "Name" | Value: "string"
type IngredientsType = map[string]string

type MicroserviceStatus struct {
	Name    string `json:"name"`     /* Name of the microservice */
	RefName string `json:"ref_name"` /* Reference Name of the microservice as it is named in the server */
	Type    string `json:"type"`

	Ingredients IngredientsType `json:"ingredients"`
}

type ServiceStatus struct {
	Name    string `json:"name"`     /* Name of the service */
	RefName string `json:"ref_name"` /* Reference Name of the service as it is named in the server */

	Microservices []MicroserviceStatus `json:"microservices"`
}

type ClientService interface {
	ListServiceStatus() (*ServiceStatus, error)

	SendAction(body models.AreaScenario, actionId, userID int) (*ActionResponseStatus, error)

	// prevOutput is an array of byte because output can be raw
	TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*ReactionResponseStatus, error)

	//
}
