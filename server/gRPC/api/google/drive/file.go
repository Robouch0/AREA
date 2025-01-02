//
// EPITECH PROJECT, 2024
// AREA
// File description:
// file

package drive

type DriveFile struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	OriginalFilename string `json:"originalFilename,omitempty"`
	Kind             string `json:"kind,omitempty"`

	Description    string `json:"description,omitempty"`
	Sha1Checksum   string `json:"sha1checksum,omitempty"`
	Sha256Checksum string `json:"sha256Checksum,omitempty"`
}
