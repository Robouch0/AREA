//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trelloClient
//

package trello_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"errors"
	"google.golang.org/grpc"
	"log"
)

type TrelloClient struct {
	MicroservicesLauncher *IServ.ReactionLauncher
	cc                    gRPCService.TrelloServiceClient
}

func NewTrelloClient(conn *grpc.ClientConn) *TrelloClient {
	micros := &IServ.ReactionLauncher{}
	trelloCli := &TrelloClient{MicroservicesLauncher: micros, cc: gRPCService.NewTrelloServiceClient(conn)}

	(*trelloCli.MicroservicesLauncher)["createBoard"] = trelloCli.createBoard
	return trelloCli
}

func (trelloCli *TrelloClient) SendAction(body models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
	return nil, nil
}

func (trelloCli *TrelloClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	return nil, nil
}

func (trelloCli *TrelloClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*trelloCli.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	log.Println(microservice)
	return nil, errors.New("No such microservice")
}

func (trelloCli *TrelloClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, action_id int) (*IServ.WebHookResponseStatus, error) {
	return nil, nil
}
