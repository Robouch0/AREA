//
// EPITECH PROJECT, 2025
// AREA
// File description:
// githubServiceTrigger
//

package github

import (
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (git *GithubService) TriggerWebHook(ctx context.Context, req *gRPCService.GithubWebHookTriggerReq) (*gRPCService.GithubWebHookTriggerReq, error) {
	act, err := git.GithubDb.GetGithubByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such action with id %d", req.ActionId)
	}

	var gitHeader http.Header
	if json.Unmarshal(req.Header, &gitHeader) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	log.Println(gitHeader)
	if act.Activated && gitHeader.Get("X-GitHub-Event") == string(act.RepoAction) {
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
