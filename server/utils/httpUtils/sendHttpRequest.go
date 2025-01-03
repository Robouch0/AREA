//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sendHttpRequest
//

package http_utils

import (
	"io"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendHttpRequest(httpRequest *http.Request, successCode int) (*http.Response, error) {
	cli := &http.Client{}
	resp, err := cli.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != successCode {
		io.Copy(os.Stderr, resp.Body)
		return nil, status.Errorf(codes.Aborted, resp.Status)
	}
	return resp, nil
}
