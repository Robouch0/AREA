//
// EPITECH PROJECT, 2025
// AREA
// File description:
// spotifyServiceInfos
//

package spotify_server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	http_utils "area/utils/httpUtils"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	spotifyMeAPIURL = "https://api.spotify.com/v1/me/player"
	spotifyArtistAPIURL = "https://api.spotify.com/v1/artists/%v"
)

type UserIDResponse struct {
	ID string `json:"id"`
}

type PlaylistResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PlaylistsResponse struct {
	Items []PlaylistResponse `json:"items"`
}

func (spot *SpotifyService) GetUserID(bearerTok string) (string, error) {
	url := "https://api.spotify.com/v1/me"
	getRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error in GETUSER id", err)
		return "", err
	}
	getRequest.Header = http.Header{}
	getRequest.Header.Set("Authorization", bearerTok)
	getRequest.Header.Set("Content-Type", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(getRequest)
	if err != nil {
		log.Println("Error in GETUSER id sending request", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error in GETUSER id status response", err)
		return "", errors.New(resp.Status)
	}

	userIDResponse := &UserIDResponse{}
	err = json.NewDecoder(resp.Body).Decode(userIDResponse)
	if err != nil {
		log.Println("Error when decoding userID res", err)
		return "", err
	}

	return userIDResponse.ID, nil
}

func (spot *SpotifyService) GetUserPlaylists(bearerTok string) (PlaylistsResponse, error) {
	url := "https://api.spotify.com/v1/me/playlists"
	getRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error in GETUSER playlists", err)
		return PlaylistsResponse{}, err
	}
	getRequest.Header = http.Header{}
	getRequest.Header.Set("Authorization", bearerTok)
	getRequest.Header.Set("Content-Type", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(getRequest)
	if err != nil {
		log.Println("Error in GETUSER playlists sending request", err)
		return PlaylistsResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var playlistsResponse PlaylistsResponse
	if err := json.Unmarshal(body, &playlistsResponse); err != nil {
		return PlaylistsResponse{}, err
	}

	for _, playlist := range playlistsResponse.Items {
		fmt.Printf("ID: %s, Name: %s\n", playlist.ID, playlist.Name)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error in GETUSER playlists status response", err)
		return PlaylistsResponse{}, errors.New(resp.Status)
	}
	return playlistsResponse, nil
}

func (spot *SpotifyService) FindPlayList(playlists PlaylistsResponse, str string) (PlaylistResponse, error) {
	for _, playlist := range playlists.Items {
		if playlist.Name == str {
			return playlist, nil
		}
	}
	return PlaylistResponse{}, errors.New("playlist not found")
}

func GetArtistInfos(ArtistId string) (*SpotifyArtistAPIResponseBody, error) {
	Url := fmt.Sprintf(spotifyArtistAPIURL, ArtistId)
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		log.Println("Request creation error: ", err)
		return nil, status.Errorf(codes.Internal, "Could not create the request: %v", err)
	}
	resp, err := http_utils.SendHttpRequest(req, 200)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var spotifyInfo SpotifyArtistAPIResponseBody
	err = json.Unmarshal(bytesBody, &spotifyInfo)
	if err != nil {
		return nil, err
	}
	return &spotifyInfo, nil
}

func GetSpotifyInfos() (*SpotifyInfoAPIResponseBody, error) {
	req, err := http.NewRequest("GET", spotifyMeAPIURL, nil)
	if err != nil {
		log.Println("Request creation error: ", err)
		return nil, status.Errorf(codes.Internal, "Could not create the request: %v", err)
	}
	resp, err := http_utils.SendHttpRequest(req, 200)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var spotifyInfo SpotifyInfoAPIResponseBody
	err = json.Unmarshal(bytesBody, &spotifyInfo)
	if err != nil {
		return nil, err
	}
	return &spotifyInfo, nil
}
