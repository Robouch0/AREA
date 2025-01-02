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
	MicroservicesLauncher *IServ.ReactionLauncher
	ActionLauncher        *IServ.ActionLauncher

	cc gRPCService.GitlabServiceClient
}

func NewGitlabClient(conn *grpc.ClientConn) *GitlabClient {
	micros := &IServ.ReactionLauncher{}
	actions := &IServ.ActionLauncher{}
	gitlab := &GitlabClient{MicroservicesLauncher: micros, ActionLauncher: actions, cc: gRPCService.NewGitlabServiceClient(conn)}
	return gitlab
}

func (git *GitlabClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No microservice Action yet")
}

func (git *GitlabClient) TriggerReaction(
	ingredients map[string]any,
	microservice string,
	prevOutput []byte,
	userID int,
) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No microservice Reaction yet")
}

func (git *GitlabClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	return nil, errors.New("No microservice TriggerWebhook yet")
}

func (git *GitlabClient) DeactivateArea(id, userID int) (*IServ.DeactivateResponseStatus, error) {
	return nil, nil
}
