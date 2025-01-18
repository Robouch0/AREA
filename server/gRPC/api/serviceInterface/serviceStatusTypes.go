//
// EPITECH PROJECT, 2024
// AREA
// File description:
// serviceStatusTypes
//

package serviceinterface

// Description of an ingredient
type IngredientDescriptor struct {
	Value       any    `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

// Map describing the possible value of an ingredient
type Ingredients = map[string]IngredientDescriptor

type MicroserviceDescriptor struct {
	Name    string `json:"name"`     /* Name of the microservice */
	RefName string `json:"ref_name"` /* Reference Name of the microservice as it is named in the server */
	Type    string `json:"type"`     /* Type of service action or reaction */

	Ingredients       Ingredients `json:"ingredients"`
	PipelineAvailable []string    `json:"pipeline_available,omitempty"`
}

type ServiceStatus struct {
	Name    string `json:"name"`     /* Name of the service */
	RefName string `json:"ref_name"` /* Reference Name of the service as it is named in the server */

	Microservices []MicroserviceDescriptor `json:"microservices"`
}
