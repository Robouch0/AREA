//
// EPITECH PROJECT, 2024
// AREA
// File description:
// discordClient
//

package discord_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
)

type DiscordClient struct {
	MicroservicesLauncher *IServ.ReactionLauncher
	cc                    gRPCService.DiscordServiceClient
}

func NewDiscordClient(conn *grpc.ClientConn) *DiscordClient {
	micros := &IServ.ReactionLauncher{}
	disCli := &DiscordClient{MicroservicesLauncher: micros, cc: gRPCService.NewDiscordServiceClient(conn)}
	(*disCli.MicroservicesLauncher)["createMsg"] = disCli.createMessage
	(*disCli.MicroservicesLauncher)["editMsg"] = disCli.editMessage
	(*disCli.MicroservicesLauncher)["deleteMsg"] = disCli.deleteMessage
	(*disCli.MicroservicesLauncher)["createReact"] = disCli.createReaction
	(*disCli.MicroservicesLauncher)["deleteAllreacts"] = disCli.deleteAllReactions
	(*disCli.MicroservicesLauncher)["deleteReact"] = disCli.deleteReaction
	return disCli
}

func (disCli *DiscordClient) createMessage(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.CreateMsg
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := disCli.cc.CreateMessage(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Content}, nil
}

func (disCli *DiscordClient) editMessage(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.EditMsg
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := disCli.cc.EditMessage(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Content}, nil
}

func (disCli *DiscordClient) deleteMessage(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.DeleteMsg
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := disCli.cc.DeleteMessage(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.MessageId}, nil
}

func (disCli *DiscordClient) createReaction(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.CreateReact
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := disCli.cc.CreateReaction(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Emoji}, nil
}

func (disCli *DiscordClient) deleteAllReactions(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.DeleteAllReact
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := disCli.cc.DeleteAllReactions(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.MessageId}, nil
}

func (disCli *DiscordClient) deleteReaction(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.DeleteReact
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := disCli.cc.DeleteReactions(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Emoji}, nil
}

func (disCli *DiscordClient) SendAction(scenario models.AreaScenario, actionId int, userID int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No action supported in Discord service (Next will be Webhooks)")
}

func (disCli *DiscordClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*disCli.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	return nil, errors.New("No such microservice")
}

func (_ *DiscordClient) TriggerWebhook(webhook *IServ.WebhookInfos, _ string, _ int) (*IServ.WebHookResponseStatus, error) {
	return &IServ.WebHookResponseStatus{}, nil
}

func (disCli *DiscordClient) DeactivateArea(id, userID int) (*IServ.DeactivateResponseStatus, error) {
	return nil, nil
}
