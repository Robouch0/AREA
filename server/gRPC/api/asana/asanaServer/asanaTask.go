//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana Task service
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
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (asana *AsanaService) CreateTask(ctx context.Context, req *gRPCService.CreateTaskReq) (*gRPCService.CreateTaskResp, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, asana.tokenDb, "AsanaService", "asana")
	if err != nil {
		return nil, err
	}

	// not checking here if the complete boolean is present atm, may cause issue later
	if req.TaskName == "" || req.ProjectName == "" || req.TaskDescription == "" || req.DueOn == "" {
		return nil, errors.New("missing some required parameters")
	}

	list, err := asana_get.ListAllProjects(tokenInfo.AccessToken)
	if err != nil {
		return nil, err
	}

	projectGid, err := asana_get.GetGidByProject(req.ProjectName, list)
	if err != nil {
		return nil, err
	}

	dueOn := strings.Split(req.DueOn, "T")[0]

	res, err := asana_create.CreateTask(tokenInfo.AccessToken, &asana_generics.AsanaBaseBody[asana_create.CreateTaskData]{
		Data: asana_create.CreateTaskData{
			Name:      req.TaskName,
			Notes:     req.TaskDescription,
			Projects:  []string{projectGid},
			Completed: req.Completion,
			DueOn:     dueOn,
		},
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error when calling Create Task: %v", err)
	}

	log.Println(res)
	return &gRPCService.CreateTaskResp{
		ProjectName:     req.ProjectName,
		TaskName:        req.TaskName,
		TaskDescription: req.TaskDescription,
		Completion:      req.Completion,
		DueOn:           req.DueOn,
	}, nil
}
