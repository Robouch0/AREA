//
// EPITECH PROJECT, 2024
// AREA
// File description:
// serviceStatusTypes
//

package serviceinterface

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
