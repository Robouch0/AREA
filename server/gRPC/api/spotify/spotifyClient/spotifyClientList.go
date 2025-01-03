//
// EPITECH PROJECT, 2025
// AREA
// File description:
// spotifyClientList
//

package spotify_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (spot *SpotifyClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Spotify",
		RefName: "spotify",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Stop the current song playing on the last device connected",
				RefName: "stopSong",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{},
			},
			{
				Name:    "Create a spotify playlist",
				RefName: "createPlaylist",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"playlistName": {
						Value:       "",
						Type:        "string",
						Description: "Name of the playlist",
						Required:    true,
					},
					"playlistDescription": {
						Value:       "",
						Type:        "string",
						Description: "Description of the playlist",
						Required:    true,
					},
					"public": {
						Value:       "",
						Type:        "string",
						Description: "Is the playlist public or private",
						Required:    true,
					},
				},
			},
			{
				Name:    "Launch the next song",
				RefName: "nextSong",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{},
			},
			{
				Name:    "Launch the previous song",
				RefName: "previousSong",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{},
			},
			{
				Name:    "Change the playback Volume",
				RefName: "setPlaybackVolume",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"volume": {
						Value:       "",
						Type:        "string",
						Description: "New volume for the song",
						Required:    true,
					},
				},
			},
			{
				Name:    "Launch a specific track",
				RefName: "launchSong",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"songUrl": {
						Value:       "",
						Type:        "string",
						Description: "URL of the song to launch",
						Required:    true,
					},
					"millisecondsPosition": {
						Value:       "",
						Type:        "string",
						Description: "Delay for the song",
						Required:    true,
					},
				},
			},
			{
				Name:    "Add a song to a playlist",
				RefName: "addSongToPlaylist",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{},
			},
		},
	}
	return status, nil
}
