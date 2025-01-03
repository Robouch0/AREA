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
	"context"
	"encoding/json"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoogleClient struct {
	MicroservicesLauncher *IServ.ReactionLauncher
	ActionLauncher        *IServ.ActionLauncher

	cc gRPCService.GoogleServiceClient
}

func NewGoogleClient(conn *grpc.ClientConn) *GoogleClient {
	micros := &IServ.ReactionLauncher{}
	actions := &IServ.ActionLauncher{}
	google := &GoogleClient{MicroservicesLauncher: micros, ActionLauncher: actions, cc: gRPCService.NewGoogleServiceClient(conn)}

	(*google.MicroservicesLauncher)["gmail/sendEmailMe"] = google.sendEmailMe
	(*google.MicroservicesLauncher)["gmail/deleteEmailMe"] = google.deleteEmailMe
	(*google.MicroservicesLauncher)["gmail/moveToTrash"] = google.moveToTrash
	(*google.MicroservicesLauncher)["gmail/moveFromTrash"] = google.moveFromTrash
	(*google.MicroservicesLauncher)["gmail/createLabel"] = google.createLabel
	(*google.MicroservicesLauncher)["gmail/deleteLabel"] = google.deleteLabel
	(*google.MicroservicesLauncher)["gmail/updateLabel"] = google.updateLabel
	// (*google.MicroservicesLauncher)["gmail/deleteDrive"] = google.deleteDrive // Not free
	// (*google.MicroservicesLauncher)["gmail/updateDrive"] = google.updateDrive
	// (*google.MicroservicesLauncher)["gmail/createDrive"] = google.createDrive
	(*google.MicroservicesLauncher)["drive/createComment"] = google.createComment
	(*google.MicroservicesLauncher)["drive/deleteComment"] = google.deleteComment
	(*google.MicroservicesLauncher)["drive/updateComment"] = google.updateComment
	(*google.MicroservicesLauncher)["drive/createEmptyFile"] = google.createEmptyFile
	(*google.MicroservicesLauncher)["drive/deleteFile"] = google.deleteFile
	(*google.MicroservicesLauncher)["drive/updateFile"] = google.updateFile
	(*google.MicroservicesLauncher)["drive/copyFile"] = google.copyFile

	(*google.ActionLauncher)["gmail/watchme"] = google.watchEmail
	(*google.ActionLauncher)["drive/watchFile"] = google.watchFile
	(*google.ActionLauncher)["drive/watchChanges"] = google.watchChanges
	return google
}

func (google *GoogleClient) watchFile(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	jsonString, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}
	var fileReq gRPCService.WatchFileReq
	err = json.Unmarshal(jsonString, &fileReq)
	if err != nil {
		return nil, err
	}
	fileReq.ActionId = uint32(actionID)
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = google.cc.WatchDriveFile(ctx, &fileReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &IServ.ActionResponseStatus{Description: "Done", ActionID: actionID}, nil
}

func (google *GoogleClient) watchChanges(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	jsonString, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}
	var fileReq gRPCService.WatchChangesReq
	err = json.Unmarshal(jsonString, &fileReq)
	if err != nil {
		return nil, err
	}
	fileReq.ActionId = uint32(actionID)
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = google.cc.WatchDriveChanges(ctx, &fileReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &IServ.ActionResponseStatus{Description: "Done", ActionID: actionID}, nil
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
	if micro, ok := (*google.ActionLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such microservice")
}

func (google *GoogleClient) deleteLabel(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var labelReq gRPCService.DeleteLabelReq
	err = json.Unmarshal(jsonString, &labelReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.DeleteLabel(ctx, &labelReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Name}, nil
}

func (google *GoogleClient) updateLabel(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var labelReq gRPCService.UpdateLabelReq
	err = json.Unmarshal(jsonString, &labelReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.UpdateLabel(ctx, &labelReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.NewName}, nil
}

func (google *GoogleClient) createLabel(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var labelReq gRPCService.CreateLabelReq
	err = json.Unmarshal(jsonString, &labelReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.CreateLabel(ctx, &labelReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Name}, nil
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

func (google *GoogleClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	// Refactor with a map
	b, err := json.Marshal(webhook.Payload)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload sent")
	}
	bHeader, err := json.Marshal(webhook.Header)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload sent")
	}
	log.Println(microservice)
	if microservice == "gmail/watchme" {
		_, err = google.cc.WatchMeTrigger(context.Background(), &gRPCService.GmailTriggerReq{Payload: b, ActionId: uint32(actionID)})
		if err != nil {
			return nil, err
		}
		return &IServ.WebHookResponseStatus{}, nil
	}
	if microservice == "drive/watchFile" {
		log.Println(microservice)
		_, err = google.cc.WatchFileTrigger(context.Background(), &gRPCService.FileTriggerReq{Headers: bHeader, ActionId: uint32(actionID)})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &IServ.WebHookResponseStatus{}, nil
	}
	if microservice == "drive/watchChanges" {
		_, err = google.cc.WatchChangesTrigger(context.Background(), &gRPCService.ChangesTriggerReq{Headers: bHeader, ActionId: uint32(actionID)})
		if err != nil {
			return nil, err
		}
		return &IServ.WebHookResponseStatus{}, nil
	}
	return nil, status.Errorf(codes.NotFound, "Microservice: %v not found", microservice)
}

func (google *GoogleClient) DeactivateArea(microservice string, id uint, userID int) (*IServ.DeactivateResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	if microservice == "gmail/watchme" {
		if _, err := google.cc.DeactivateGmailAction(ctx, &gRPCService.DeactivateGmail{ActionId: uint32(id)}); err != nil {
			return nil, err
		}
	}
	if microservice == "drive/watchChanges" || microservice == "drive/watchFile" {
		if _, err := google.cc.DeactivateDriveAction(ctx, &gRPCService.DeactivateDrive{ActionId: uint32(id)}); err != nil {
			return nil, err
		}
	}
	return &IServ.DeactivateResponseStatus{
		ActionID:    id,
		Description: "",
	}, nil
}
