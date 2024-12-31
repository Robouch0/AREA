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
	_, err := google.driveDb.StoreNewGWatch(&models.Drive{
		ActionID:  uint(req.ActionId),
		UserID:    uint(tokenInfo.UserID),
		Activated: true,
		ChannelID: channelID,
	})
	if err != nil {
		return nil, err
	}
	_, err = drive.WatchFile(tokenInfo.AccessToken, fileID, channelID, uint(req.ActionId))
	if err != nil {
		return nil, err
	}
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
	_, err = google.driveDb.StoreNewGWatch(&models.Drive{
		ActionID:  uint(req.ActionId),
		UserID:    uint(tokenInfo.UserID),
		Activated: true,
		ChannelID: channelID,
	})
	if err != nil {
		return nil, err
	}
	_, err = drive.WatchChanges(tokenInfo.AccessToken, channelID, uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (google *GoogleService) WatchFileTrigger(ctx context.Context, req *gRPCService.FileTriggerReq) (*gRPCService.FileTriggerReq, error) {
	// Get type for drive payload
	var header http.Header
	if json.Unmarshal(req.Headers, &header) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}

	actionID := header.Get("X-Goog-Channel-ID")
	act, err := google.driveDb.GetByChannelID(actionID)
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
	// Get type for drive payload
	var header http.Header
	if json.Unmarshal(req.Headers, &header) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}

	actionID := header.Get("X-Goog-Channel-ID")
	act, err := google.driveDb.GetByChannelID(actionID)
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
