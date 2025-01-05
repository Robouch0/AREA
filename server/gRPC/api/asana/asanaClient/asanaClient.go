//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asanaClient
//

package asana_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"errors"
	"google.golang.org/grpc"
	"log"
)

type AsanaClient struct {
	MicroservicesLauncher *IServ.ReactionLauncher
	cc                    gRPCService.AsanaServiceClient
}

func NewAsanaClient(conn *grpc.ClientConn) *AsanaClient {
	micros := &IServ.ReactionLauncher{}
	asana := &AsanaClient{MicroservicesLauncher: micros, cc: gRPCService.NewAsanaServiceClient(conn)}

	(*asana.MicroservicesLauncher)["createProject"] = asana.createProject
	return asana
}

func (asana *AsanaClient) SendAction(body models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
	return nil, nil
}

func (asana *AsanaClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	return nil, nil
}

func (asana *AsanaClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*asana.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	log.Println(microservice)
	return nil, errors.New("no such microservice")
}

func (asana *AsanaClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, action_id int) (*IServ.WebHookResponseStatus, error) {
	return nil, nil
}
