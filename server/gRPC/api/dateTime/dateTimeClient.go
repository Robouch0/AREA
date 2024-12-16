//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTimeClient
//

package dateTime

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
)

type DTServiceClient struct {
	gRPCService.DateTimeServiceClient
}

func NewDateTimeServiceClient(conn *grpc.ClientConn) *DTServiceClient {
	return &DTServiceClient{gRPCService.NewDateTimeServiceClient(conn)}
}

func (git *DTServiceClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Date and Time API",
		RefName: "dt",

		Microservices: []IServ.MicroserviceStatus{
			IServ.MicroserviceStatus{
				Name:    "Trigger a reaction at a specific date and time",
				RefName: "timeTrigger",
				Type:    "action",

				Ingredients: map[string]string{
					// "activated": "bool",
					"minutes":   "int",
					"hours":     "int",
					"day_month": "int",
					"month":     "int",
					"day_week":  "int",
				},
			},
		},
	}
	return status, nil
}

func (react *DTServiceClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No reaction available for this service")
}

func (dt *DTServiceClient) SendAction(body map[string]any, actionID int) (*IServ.ActionResponseStatus, error) {
	jsonString, err := json.Marshal(body["action"])
	if err != nil {
		return nil, err
	}

	action := models.Action{}
	err = json.Unmarshal(jsonString, &action)
	if err != nil {
		return nil, err
	}

	timeReqJson, err := json.Marshal(action.Ingredients)
	if err != nil {
		return nil, err
	}

	timeReq := gRPCService.TriggerTimeRequest{ActionId: int32(actionID)}
	err = json.Unmarshal(timeReqJson, &timeReq)
	if err != nil {
		return nil, err
	}
	dt.LaunchCronJob(context.Background(), &timeReq)
	return &IServ.ActionResponseStatus{Description: "Done", ActionID: actionID}, nil
}
