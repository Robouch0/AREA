//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana Board service
//

package asana_server

import (
	"area/gRPC/api/asana/asanaCreate"
	asana_generics "area/gRPC/api/asana/asanaGenerics"
	"area/gRPC/api/asana/asanaGet"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (asana *AsanaService) CreateProject(ctx context.Context, req *gRPCService.CreateProjectReq) (*gRPCService.CreateProjectResp, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, asana.tokenDb, "AsanaService", "asana")
	if err != nil {
		return nil, err
	}

	if req.ProjectName == "" || req.Color == "" || req.DefaultView == "" || req.WorkspaceName == "" {
		return nil, errors.New("project name color and default view required atleast")
	}

	list, err := asana_get.ListWorkspace(tokenInfo.AccessToken)

	if err != nil {
		return nil, err
	}

	workspaceGid, err := asana_get.GetGidByWorkspace(req.WorkspaceName, list)

	if err != nil {
		return nil, err
	}

	res, err := asana_create.CreateProject(tokenInfo.AccessToken, &asana_generics.AsanaBaseBody[asana_create.CreateProjectData]{
		Data: asana_create.CreateProjectData{
			Name:         req.ProjectName,
			Color:        req.Color,
			DefaultView:  req.DefaultView,
			WorkspaceGid: workspaceGid,
		},
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error when calling Create Project : %v", err)
	}
	log.Println(res)
	return &gRPCService.CreateProjectResp{}, nil
}
