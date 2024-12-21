//
// EPITECH PROJECT, 2024
// AREA
// File description:
// deleteEmail
//

package gmail

import (
	"area/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	deleteURL = "https://gmail.googleapis.com/gmail/v1/users/%s/messages/%s"
)

func DeleteEmail(googleUserID string, accessToken string, messageID string) error {
	url := fmt.Sprintf(deleteURL, googleUserID, messageID)

	request, _ := http.NewRequest("DELETE", url, nil) // GetUserMail request
	request.Header = utils.GetDefaultBearerHTTPHeader(accessToken)
	request.Header.Add("Accept", "application/json")

	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		return err
	}
	if result.StatusCode != 200 {
		io.Copy(os.Stderr, result.Body)
		return status.Errorf(codes.Aborted, result.Status)
	}
	log.Println("Message with id: ", messageID, " deleted.")
	return nil
}
