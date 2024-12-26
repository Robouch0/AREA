//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleClient
//

package google_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoogleClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	ActionLauncher        *IServ.ActionLauncher

	cc gRPCService.GoogleServiceClient
}

func NewGoogleClient(conn *grpc.ClientConn) *GoogleClient {
	micros := &IServ.MicroserviceLauncher{}
	actions := &IServ.ActionLauncher{}
	google := &GoogleClient{MicroservicesLauncher: micros, ActionLauncher: actions, cc: gRPCService.NewGoogleServiceClient(conn)}
	(*google.MicroservicesLauncher)["gmail/sendEmailMe"] = google.sendEmailMe
	(*google.MicroservicesLauncher)["gmail/deleteEmailMe"] = google.deleteEmailMe
	(*google.MicroservicesLauncher)["gmail/moveToTrash"] = google.moveToTrash
	(*google.MicroservicesLauncher)["gmail/moveFromTrash"] = google.moveFromTrash

	(*google.ActionLauncher)["gmail/watchme"] = google.watchEmail
	return google
}

func (google *GoogleClient) watchEmail(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := google.cc.WatchGmailEmail(ctx, &gRPCService.EmailTriggerReq{ActionId: uint32(actionID)})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &IServ.ActionResponseStatus{Description: "Done", ActionID: actionID}, nil
}

func (google *GoogleClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*google.ActionLauncher)[scenario.Reaction.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such microservice")
}

func (google *GoogleClient) moveToTrash(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var deleteEmailMe gRPCService.TrashEmailRequestMe
	err = json.Unmarshal(jsonString, &deleteEmailMe)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.MoveToTrash(ctx, &deleteEmailMe)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Subject}, nil
}

func (google *GoogleClient) moveFromTrash(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var deleteEmailMe gRPCService.TrashEmailRequestMe
	err = json.Unmarshal(jsonString, &deleteEmailMe)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.MoveFromTrash(ctx, &deleteEmailMe)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Subject}, nil
}

func (google *GoogleClient) deleteEmailMe(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var deleteEmailMe gRPCService.DeleteEmailRequestMe
	err = json.Unmarshal(jsonString, &deleteEmailMe)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.DeleteEmailMe(ctx, &deleteEmailMe)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Subject}, nil
}

func (google *GoogleClient) sendEmailMe(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var sendEmailMe gRPCService.EmailRequestMe
	err = json.Unmarshal(jsonString, &sendEmailMe)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.SendEmailMe(ctx, &sendEmailMe)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.BodyMessage}, nil
}

func (google *GoogleClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*google.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	return nil, errors.New("No such microservice")
}

func (_ *GoogleClient) TriggerWebhook(_ map[string]any, _ string, _ int) (*IServ.WebHookResponseStatus, error) {
	return &IServ.WebHookResponseStatus{}, nil
}
