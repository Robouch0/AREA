//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleDriveWebhook
//

package google_server

import (
	"area/gRPC/api/google/drive"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func watchFile(
	google *GoogleService,
	req *gRPCService.WatchFileReq,
	tokenInfo *models.Token,
	fileID string,
) (*gRPCService.WatchFileReq, error) {
	channelID := uuid.NewString()
	respWatch, err := drive.WatchFile(tokenInfo.AccessToken, fileID, channelID, uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	_, err = google.driveDb.StoreNewGWatch(&models.Drive{
		ActionID:   uint(req.ActionId),
		UserID:     uint(tokenInfo.UserID),
		Activated:  true,
		ChannelID:  channelID,
		ResourceID: respWatch.ResourceId,
		FileName:   req.FileName,
	})
	if err != nil {
		return nil, err
	}
	log.Println("New WatchFile for drive")
	return req, nil
}

func (google *GoogleService) WatchDriveFile(ctx context.Context, req *gRPCService.WatchFileReq) (*gRPCService.WatchFileReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	list, err := drive.ListFiles(tokenInfo.AccessToken)
	if err != nil {
		return nil, err
	}
	for _, f := range list.Files {
		if f.Name == req.FileName {
			return watchFile(google, req, tokenInfo, f.ID)
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "No such file: %v", req.FileName)
}

func (google *GoogleService) WatchDriveChanges(ctx context.Context, req *gRPCService.WatchChangesReq) (*gRPCService.WatchChangesReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	channelID := uuid.NewString()
	respWatch, err := drive.WatchChanges(tokenInfo.AccessToken, channelID, uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	_, err = google.driveDb.StoreNewGWatch(&models.Drive{
		ActionID:   uint(req.ActionId),
		UserID:     uint(tokenInfo.UserID),
		Activated:  true,
		ChannelID:  channelID,
		ResourceID: respWatch.ResourceId,
		FileName:   "",
	})
	if err != nil {
		return nil, err
	}
	log.Println("New WatchChanges for drive")
	return req, nil
}

func (google *GoogleService) WatchFileTrigger(ctx context.Context, req *gRPCService.FileTriggerReq) (*gRPCService.FileTriggerReq, error) {
	var header http.Header
	if json.Unmarshal(req.Headers, &header) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}

	act, err := google.driveDb.GetActionByID(strconv.Itoa(int(req.ActionId)))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if act.Activated {
		ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err := google.reactService.LaunchReaction(
			ctx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: []byte(req.Headers)},
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return req, nil
}

func (google *GoogleService) WatchChangesTrigger(ctx context.Context, req *gRPCService.ChangesTriggerReq) (*gRPCService.ChangesTriggerReq, error) {
	var header http.Header
	if json.Unmarshal(req.Headers, &header) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}

	act, err := google.driveDb.GetActionByID(strconv.Itoa(int(req.ActionId)))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if act.Activated {
		ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err := google.reactService.LaunchReaction(
			ctx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: []byte(req.Headers)},
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return req, nil
}

func (google *GoogleService) SetActivateDriveAction(ctx context.Context, req *gRPCService.SetActivateDrive) (*gRPCService.SetActivateDrive, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleGmailService", "google")
	if err != nil {
		return nil, err
	}
	act, err := google.driveDb.GetUserActionByID(uint(tokenInfo.UserID), uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	if !req.Activated {
		err = drive.StopWatchDrive(tokenInfo, act.ChannelID, act.ResourceID)
	} else {
		if req.Microservice == "watchChanges" {
			_, err = drive.WatchChanges(tokenInfo.AccessToken, act.ChannelID, act.ActionID)
		} else {
			_, err = drive.WatchFile(tokenInfo.AccessToken, act.FileName, act.ChannelID, act.ActionID)
		}
	}
	if err != nil {
		return nil, err
	}
	_, err = google.gmailDb.SetActivateByActionID(req.Activated, uint(tokenInfo.UserID), uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}
