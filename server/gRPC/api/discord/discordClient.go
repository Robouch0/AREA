//
// EPITECH PROJECT, 2024
// AREA
// File description:
// discordClient
//

package discord

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
)

type DiscordClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	cc                    gRPCService.DiscordServiceClient
}

func NewDiscordClient(conn *grpc.ClientConn) *DiscordClient {
	micros := &IServ.MicroserviceLauncher{}
	disCli := &DiscordClient{MicroservicesLauncher: micros, cc: gRPCService.NewDiscordServiceClient(conn)}
	(*disCli.MicroservicesLauncher)["createMsg"] = disCli.CreateMessage
	(*disCli.MicroservicesLauncher)["createReact"] = disCli.CreateReaction
	return disCli
}

func (disCli *DiscordClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Discord",
		RefName: "discord",

		Microservices: []IServ.MicroserviceStatus{
			IServ.MicroserviceStatus{
				Name:    "Create a message in a channel",
				RefName: "createMsg",
				Type:    "reaction",

				Ingredients: map[string]string{
					"channel": "string",
					"content": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Edit a message",
				RefName: "editMsg",
				Type:    "reaction",

				Ingredients: map[string]string{
					"channel":   "string",
					"messageid": "string",
					"content":   "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Delete a message",
				RefName: "deleteMsg",
				Type:    "reaction",

				Ingredients: map[string]string{
					"channel":   "string",
					"messageid": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Create a reaction on a message",
				RefName: "createReact",
				Type:    "reaction",

				Ingredients: map[string]string{
					"channel":   "string",
					"messageid": "string",
					"emoji":     "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Delete all reactions on a message",
				RefName: "deleteAllreacts",
				Type:    "reaction",

				Ingredients: map[string]string{
					"channel":   "string",
					"messageid": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Delete selected reactions on a message",
				RefName: "deleteReact",
				Type:    "reaction",

				Ingredients: map[string]string{
					"channel":   "string",
					"messageid": "string",
					"emoji":     "string",
				},
			},
		},
	}
	return status, nil
}

func (disCli *DiscordClient) CreateMessage(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.CreateMsg
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	res, err := disCli.cc.CreateMessage(context.Background(), &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Content}, nil
}

func (disCli *DiscordClient) EditMessage(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.EditMsg
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	res, err := disCli.cc.EditMessage(context.Background(), &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Content}, nil
}

func (disCli *DiscordClient) DeleteMessage(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.DeleteMsg
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	res, err := disCli.cc.DeleteMessage(context.Background(), &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.MessageId}, nil
}

func (disCli *DiscordClient) CreateReaction(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.CreateReact
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	res, err := disCli.cc.CreateReaction(context.Background(), &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Emoji}, nil
}

func (disCli *DiscordClient) deleteAllReactions(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.DeleteAllReact
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	res, err := disCli.cc.DeleteAllReactions(context.Background(), &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.MessageId}, nil
}

func (disCli *DiscordClient) DeleteReaction(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.DeleteReact
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	res, err := disCli.cc.DeleteReactions(context.Background(), &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Emoji}, nil
}

func (disCli *DiscordClient) SendAction(body map[string]any, actionId int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No action supported in Discord service (Next will be Webhooks)")
}

func (disCli *DiscordClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*disCli.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput)
	}
	return nil, errors.New("No such microservice")
}
