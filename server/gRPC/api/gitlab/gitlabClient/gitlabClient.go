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

//https://www.gitlab.com/api/v4/

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

func (git *GitlabClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*git.ActionLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such action microservice")
}

func (git *GitlabClient) TriggerReaction(
	ingredients map[string]any,
	microservice string,
	prevOutput []byte,
	userID int,
) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*git.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	return nil, errors.New("No such microservice")
}

func (git *GitlabClient) TriggerWebhook(payload map[string]any, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	return nil, errors.New("No microservice TriggerWebhook yet")
}
