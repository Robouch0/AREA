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
				Name:    "Check every 3 minutes if an artist has a certain number of followers",
				RefName: "checkFollowers",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{
					"artist_id": {
						Value:       "",
						Type:        "string",
						Description: "Spotify Id of the artist",
						Required:    true,
					},
					"followers": {
						Value:       "",
						Type:        "int",
						Description: "numbers of follow up to check",
						Required:    true,
					},
				},
				PipelineAvailable: []string{},
			},
			{
				Name:    "Check every 1 minutes if the volume exceed a certain amount",
				RefName: "checkVolume",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{
					"volume": {
						Value:       "",
						Type:        "int",
						Description: "volume threshold",
						Required:    true,
					},
				},
				PipelineAvailable: []string{},
			},
			{
				Name:    "Check every 3 minutes if the song is on repeat",
				RefName: "checkRepeat",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{},
			},
			{
				Name:    "Check every 3 minutes if the playlist is shuffle",
				RefName: "checkShuffle",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{},
			},
			{
				Name:    "Check every 3 minutes if the song is playing",
				RefName: "checkPlaying",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{},
			},
			{
				Name:    "Stop the current song playing on the last device connected",
				RefName: "stopSong",
				Type:    "reaction",

				Ingredients:       map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{},
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
				PipelineAvailable: []string{"playlistName", "playlistDescription", "public"},
			},
			{
				Name:    "Launch the next song",
				RefName: "nextSong",
				Type:    "reaction",

				Ingredients:       map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{},
			},
			{
				Name:    "Launch the previous song",
				RefName: "previousSong",
				Type:    "reaction",

				Ingredients:       map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{},
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
				PipelineAvailable: []string{"volume"},
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
				PipelineAvailable: []string{"songUrl", "millisecondsPosition"},
			},
			//{
			//	Name:    "Add a song to a playlist",
			//	RefName: "addSongToPlaylist",
			//	Type:    "reaction",
			//
			//	Ingredients: map[string]IServ.IngredientDescriptor{},
			//},
		},
	}
	return status, nil
}
