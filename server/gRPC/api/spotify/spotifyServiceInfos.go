package spotify

import (
	"encoding/json"
	"errors"
	"net/http"
	"log"
)

type UserIDResponse struct {
	ID string `json:"id"`
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
