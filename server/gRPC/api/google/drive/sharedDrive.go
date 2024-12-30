//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sharedDrive
//

package drive

type SharedDrive struct {
	ID   string
	Name string
	Kind string `json:"kind"` // Available for output only

	CreatedTime string `json:"createdTime"`
	// Later add capabilities and restrictionss
}
