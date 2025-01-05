//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana Board service
//

package asana_server

import (
	asana_generics "area/gRPC/api/asana/asanaGenerics"
	asana_project "area/gRPC/api/asana/asanaProject"
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

	list, err := asana_project.ListWorkspace(tokenInfo.AccessToken)

	if err != nil {
		return nil, err
	}

	workspaceGid, err := asana_project.GetGidByWorkspace(req.WorkspaceName, list)

	if err != nil {
		return nil, err
	}

	res, err := asana_project.CreateProject(tokenInfo.AccessToken, &asana_generics.AsanaBaseBody[asana_project.CreateProjectData]{
		Data: asana_project.CreateProjectData{
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
