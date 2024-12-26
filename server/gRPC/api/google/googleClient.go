//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleClient
//

package google

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
)

type GoogleClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	cc                    gRPCService.GoogleServiceClient
}

func NewGoogleClient(conn *grpc.ClientConn) *GoogleClient {
	micros := &IServ.MicroserviceLauncher{}
	google := &GoogleClient{MicroservicesLauncher: micros, cc: gRPCService.NewGoogleServiceClient(conn)}
	(*google.MicroservicesLauncher)["gmail/sendEmailMe"] = google.sendEmailMe
	(*google.MicroservicesLauncher)["gmail/deleteEmailMe"] = google.deleteEmailMe
	(*google.MicroservicesLauncher)["gmail/moveToTrash"] = google.moveToTrash
	(*google.MicroservicesLauncher)["gmail/moveFromTrash"] = google.moveFromTrash

	return google
}

func (google *GoogleClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Google",
		RefName: "google",

		Microservices: []IServ.MicroserviceStatus{
			IServ.MicroserviceStatus{
				Name:    "Send an email to a specific user",
				RefName: "gmail/sendEmailMe",
				Type:    "reaction",

				Ingredients: map[string]string{
					"to":           "string",
					"subject":      "string",
					"body_message": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Delete an email of a specific user",
				RefName: "gmail/deleteEmailMe",
				Type:    "reaction",

				Ingredients: map[string]string{
					"subject": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Move an email to trash",
				RefName: "gmail/moveToTrash",
				Type:    "reaction",

				Ingredients: map[string]string{
					"subject": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Move an email from trash",
				RefName: "gmail/moveFromTrash",
				Type:    "reaction",

				Ingredients: map[string]string{
					"subject": "string",
				},
			},
		},
	}
	return status, nil
}

func (google *GoogleClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No action supported in hugging face service (Next will be Webhooks)")
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

func (_ *GoogleClient) TriggerWebhook(_ map[string]any, _ string, _ int) {

}
