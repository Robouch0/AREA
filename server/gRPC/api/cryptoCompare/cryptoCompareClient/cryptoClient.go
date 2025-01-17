//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// cryptoClient
//

package cryptocompareclient

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"errors"

	"google.golang.org/grpc"
)

type CryptoClient struct {
	ActionLauncher *IServ.ActionLauncher

	cc gRPCService.CryptoServiceClient
}

func NewCryptoClient(conn *grpc.ClientConn) *CryptoClient {
	actions := &IServ.ActionLauncher{}
	crypto := &CryptoClient{ActionLauncher: actions, cc: gRPCService.NewCryptoServiceClient(conn)}

	(*crypto.ActionLauncher)["cryptoExceed"] = crypto.SendIsHigherAction
	(*crypto.ActionLauncher)["cryptoLower"] = crypto.SendIsLowerAction

	// send a request to create a cron job and every 30min trigger it
	return crypto
}

func (crypto *CryptoClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*crypto.ActionLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such microservice")
}

func (crypto *CryptoClient) TriggerReaction(
	ingredients map[string]any,
	microservice string,
	userID int,
) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No microservice Reaction yet")
}

func (crypto *CryptoClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	return nil, errors.New("No microservice TriggerWebhook yet")
}

func (crypto *CryptoClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := crypto.cc.SetActivate(ctx, &gRPCService.SetActivateCrypto{
		ActionId:  uint32(id),
		Activated: activated,
	})
	if err != nil {
		return nil, err
	}
	return &IServ.SetActivatedResponseStatus{
		ActionID:    id,
		Description: "CryptoCompare Deactivated",
	}, nil
}
