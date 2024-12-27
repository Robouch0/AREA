//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gitlabClient
//

package gitlab_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"errors"

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

func (google *GitlabClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No microservice Action yet")
}

func (google *GitlabClient) TriggerReaction(
	ingredients map[string]any,
	microservice string,
	prevOutput []byte,
	userID int,
) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No microservice Reaction yet")
}

func (google *GitlabClient) TriggerWebhook(payload map[string]any, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	return nil, errors.New("No microservice TriggerWebhook yet")
}
