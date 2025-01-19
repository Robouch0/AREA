//
// EPITECH PROJECT, 2025
// AREA
// File description:
// githubServiceTrigger
//

package github

import (
	githubtypes "area/gRPC/api/github/githubTypes"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"net/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PipelinePayload struct {
	RepoOwner string `json:"owner"`
	RepoName string  `json:"repo"`
}

func (git *GithubService) TriggerWebHook(ctx context.Context, req *gRPCService.GithubWebHookTriggerReq) (*gRPCService.GithubWebHookTriggerReq, error) {
	act, err := git.GithubDb.GetGithubByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such action with id %d", req.ActionId)
	}

	var gitHeader http.Header
	if json.Unmarshal(req.Header, &gitHeader) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	var gitpayload githubtypes.GithubWebHookResponse
	if json.Unmarshal(req.Payload, &gitpayload) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	if act.Activated && gitHeader.Get("X-GitHub-Event") == string(act.RepoAction) && gitpayload.Action == string(act.ActionType){
		Prepayload := &PipelinePayload{RepoOwner: act.RepoOwner, RepoName: act.RepoName}
		payload, err := json.Marshal(Prepayload)
		if err != nil {
			return nil, err
		}
		reactCtx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err = git.reactService.LaunchReaction(
			reactCtx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: []byte(payload)},
		)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not handle action's reaction")
		}
	}
	return req, nil
}
