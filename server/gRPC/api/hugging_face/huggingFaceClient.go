//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceClient
//

package huggingFace

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"errors"

	"google.golang.org/grpc"
)

type HuggingFaceServiceClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	cc                    gRPCService.HuggingFaceServiceClient
}

func NewHuggingFaceClient(conn *grpc.ClientConn) *HuggingFaceServiceClient {
	micros := &IServ.MicroserviceLauncher{}
	hfCli := &HuggingFaceServiceClient{MicroservicesLauncher: micros, cc: gRPCService.NewHuggingFaceServiceClient(conn)}
	(*hfCli.MicroservicesLauncher)["textGen"] = hfCli.sendTextGenerationReaction
	return hfCli
}

func (hfCli *HuggingFaceServiceClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Hugging Face",
		RefName: "hf",

		Microservices: []IServ.MicroserviceStatus{
			IServ.MicroserviceStatus{
				Name:    "Text Generation",
				RefName: "textGen",
				Type:    "reaction",

				Ingredients: map[string]string{
					"model":  "string",
					"inputs": "string",
				},
			},
		},
	}
	return status, nil
}

func (hfCli *HuggingFaceServiceClient) SendAction(scenario models.AreaScenario, actionId int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No action supported in hugging face service (Next will be Webhooks)")
}

func (hfCli *HuggingFaceServiceClient) sendTextGenerationReaction(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	res, err := hfCli.cc.LaunchTextGeneration(context.Background(), &gRPCService.TextGenerationReq{Inputs: ingredients["inputs"].(string)})
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.GeneratedText}, nil
}

func (hfCli *HuggingFaceServiceClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*hfCli.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput)
	}
	return nil, errors.New("No such microservice")
}
