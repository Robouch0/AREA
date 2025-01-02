//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleDriveFiles
//

package google_server

import (
	"area/gRPC/api/google/drive"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
)

func (google *GoogleService) CreateEmptyFile(ctx context.Context, req *gRPCService.CreateEmptyFileReq) (*gRPCService.CreateEmptyFileReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	_, err = drive.CreateEmptyFile(tokenInfo.AccessToken, drive.DriveFile{
		Name:        req.FileName,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (google *GoogleService) DeleteFile(ctx context.Context, req *gRPCService.DeleteFileReq) (*gRPCService.DeleteFileReq, error) {
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
			err = drive.DeleteFile(tokenInfo.AccessToken, f.ID)
			if err != nil {
				return nil, err
			}
			return req, nil
		}
	}
	return req, nil
}

func (google *GoogleService) UpdateFileMetadata(ctx context.Context, req *gRPCService.UpdateFileMetaReq) (*gRPCService.UpdateFileMetaReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	list, err := drive.ListFiles(tokenInfo.AccessToken)
	if err != nil {
		return nil, err
	}
	for _, f := range list.Files {
		if f.Name == req.OldFileName {
			_, err := drive.UpdateFile(tokenInfo.AccessToken, f.ID, drive.DriveFile{
				Name:        req.NewFileName,
				Description: req.Description,
			})
			if err != nil {
				return nil, err
			}
			return req, nil
		}
	}
	return req, nil
}

func (google *GoogleService) CopyFile(ctx context.Context, req *gRPCService.CopyFileReq) (*gRPCService.CopyFileReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	list, err := drive.ListFiles(tokenInfo.AccessToken)
	if err != nil {
		return nil, err
	}
	for _, f := range list.Files {
		if f.Name == req.SrcFileName {
			_, err := drive.CopyFile(tokenInfo.AccessToken, f.ID, drive.DriveFile{
				Name:        req.DestFileName,
				Description: req.Description,
			})
			if err != nil {
				return nil, err
			}
			return req, nil
		}
	}
	return req, nil
}
