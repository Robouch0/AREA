//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTimeClient
//

package dateTime_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
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

		Microservices: []IServ.MicroserviceDescriptor{
			IServ.MicroserviceDescriptor{
				Name:    "Trigger a reaction at a specific date and time",
				RefName: "timeTrigger",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"minutes": {
						Value:       0,
						Type:        "int",
						Description: "Minutes",
						Required:    true,
					},
					"hours": {
						Value:       0,
						Type:        "int",
						Description: "Hours",
						Required:    true,
					},
					"day_month": {
						Value:       0,
						Type:        "int",
						Description: "Day of the current month",
						Required:    true,
					},
					"month": {
						Value:       0,
						Type:        "int",
						Description: "Month number",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}

func (react *DTServiceClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No reaction available for this service")
}

func (dt *DTServiceClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	timeReqJson, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	timeReq := gRPCService.TriggerTimeRequest{ActionId: int32(actionID)}
	err = json.Unmarshal(timeReqJson, &timeReq)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	dt.LaunchCronJob(ctx, &timeReq)
	return &IServ.ActionResponseStatus{Description: "Done", ActionID: actionID}, nil
}

func (_ *DTServiceClient) TriggerWebhook(webhook *IServ.WebhookInfos, _ string, _ int) (*IServ.WebHookResponseStatus, error) {
	return &IServ.WebHookResponseStatus{}, nil
}

func (dt *DTServiceClient) DeactivateArea(microservice string, id uint, userID int) (*IServ.DeactivateResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := dt.DeactivateAction(ctx, &gRPCService.DeactivateTime{
		ActionId: uint32(id),
	})
	if err != nil {
		return nil, err
	}
	return &IServ.DeactivateResponseStatus{
		ActionID:    id,
		Description: "DateTime Deactivated",
	}, nil
}
