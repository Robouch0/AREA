//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana Board service
//

package asana_server

import (
	asana_create "area/gRPC/api/asana/asanaCreate"
	asana_generics "area/gRPC/api/asana/asanaGenerics"
	asana_get "area/gRPC/api/asana/asanaGet"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (asana *AsanaService) CreateSection(ctx context.Context, req *gRPCService.CreateSectionReq) (*gRPCService.CreateSectionResp, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, asana.tokenDb, "AsanaService", "asana")
	if err != nil {
		return nil, err
	}

	if req.SectionName == "" || req.ProjectName == "" {
		return nil, errors.New("section name and project name are required")
	}

	list, err := asana_get.ListAllProjects(tokenInfo.AccessToken)
	if err != nil {
		return nil, err
	}

	projectGid, err := asana_get.GetGidByProject(req.ProjectName, list)
	if err != nil {
		return nil, err
	}

	res, err := asana_create.CreateSection(tokenInfo.AccessToken, &asana_generics.AsanaBaseBody[asana_create.CreateSectionData]{
		Data: asana_create.CreateSectionData{
			Name:       req.SectionName,
			ProjectGid: projectGid,
		},
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error when calling Create Section: %v", err)
	}
	log.Println(res)
	return &gRPCService.CreateSectionResp{
		ProjectName: req.ProjectName,
		SectionName: req.SectionName,
	}, nil
}
