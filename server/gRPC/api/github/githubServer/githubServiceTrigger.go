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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (git *GithubService) TriggerWebHook(ctx context.Context, req *gRPCService.GithubWebHookTriggerReq) (*gRPCService.GithubWebHookTriggerReq, error) {
	act, err := git.GithubDb.GetGithubByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such action with id %d", req.ActionId)
	}

	var gitpayload githubtypes.GithubEvents
	if json.Unmarshal(req.Payload, &gitpayload) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	if len(gitpayload.Hook.Events) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "No events in github webhook payload")
	}
	if act.Activated && gitpayload.Hook.Events[0] == string(act.RepoAction) {
		reactCtx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err = git.reactService.LaunchReaction(
			reactCtx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: req.Payload},
		)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not handle action's reaction")
		}
	}
	return req, nil
}
