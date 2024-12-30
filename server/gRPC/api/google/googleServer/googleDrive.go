//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleDrive
//

package google_server

import (
	"area/gRPC/api/google/drive"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
)

func (google *GoogleService) CreateSharedDrive(ctx context.Context, req *gRPCService.CreateSharedDriveReq) (*gRPCService.CreateSharedDriveReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	_, err = drive.CreateSharedDrive(tokenInfo.AccessToken, drive.SharedDrive{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (google *GoogleService) DeleteSharedDrive(ctx context.Context, req *gRPCService.DeleteSharedDriveReq) (*gRPCService.DeleteSharedDriveReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	if err := drive.DeleteSharedDrive(tokenInfo.AccessToken, req.Name, req.UseDomainAdminAccess); err != nil {
		return nil, err
	}
	return req, nil
}

func (google *GoogleService) UpdateSharedDrive(ctx context.Context, req *gRPCService.UpdateSharedDriveReq) (*gRPCService.UpdateSharedDriveReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	_, err = drive.UpdateSharedDrive(tokenInfo.AccessToken, req.OldName, drive.SharedDrive{
		Name:   req.NewName,
		Hidden: req.Hidden,
	}, req.UseDomainAdminAccess)
	if err != nil {
		return nil, err
	}
	return req, nil
}
