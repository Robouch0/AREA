//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceService
//

package huggingFace

import (
	gRPCService "area/protogen/gRPC/proto"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type HuggingFaceService struct {
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedHuggingFaceServiceServer
}

func NewHuggingFaceService() HuggingFaceService {
	return HuggingFaceService{reactService: nil}
}

func (hfServ *HuggingFaceService) LaunchTextGeneration(_ context.Context, req *gRPCService.TextGenerationReq) (*gRPCService.TextGenerationRes, error) {
	url := "https://api-inference.huggingface.co/models/google/gemma-2-2b-it"
	bearerTok := "Bearer " + "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" // DO NOT DO THAT AT HOME os.Getenv("API_HUGGING_FACE")

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	postRequest.Header.Set("Authorization", bearerTok)
	postRequest.Header.Add("Content-Type", "application/json;charset=UTF-8")
	postRequest.Header.Add("Accept", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(postRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	textGenRes := &[]gRPCService.TextGenerationRes{}
	if err = json.NewDecoder(resp.Body).Decode(textGenRes); err != nil {
		return nil, err
	}
	if len(*textGenRes) == 0 {
		return nil, errors.New("No generated text received")
	}
	return &(*textGenRes)[0], nil
}
