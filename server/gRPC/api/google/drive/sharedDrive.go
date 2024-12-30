//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sharedDrive
//

package drive

// IT IS NOT FREE : (

type SharedDrive struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Kind        string `json:"kind,omitempty"` // Available for output only
	Hidden      bool   `json:"hidden,omitempty"`
	CreatedTime string `json:"createdTime,omitempty"`
	// Later add capabilities and restrictionss
}
