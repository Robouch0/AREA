//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gitlabClient
//

package gitlab_client

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"

	"google.golang.org/grpc"
)

type GitlabClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	ActionLauncher        *IServ.ActionLauncher

	cc gRPCService.GitlabServiceClient
}

func NewGitlabClient(conn *grpc.ClientConn) *GitlabClient {
	micros := &IServ.MicroserviceLauncher{}
	actions := &IServ.ActionLauncher{}
	gitlab := &GitlabClient{MicroservicesLauncher: micros, ActionLauncher: actions, cc: gRPCService.NewGitlabServiceClient(conn)}
	return gitlab
}
