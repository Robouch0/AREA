//
// EPITECH PROJECT, 2024
// AREA
// File description:
// githubService
//

package github

import (
	"area/db"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	"cmp"
	"fmt"

	"google.golang.org/grpc"
)

type GithubService struct {
	tokenDb      *db.TokenDb
	GithubDb     *db.GithubDB
	reactService gRPCService.ReactionServiceClient

	// Map each event to a lambda
	gRPCService.UnimplementedGithubServiceServer
}

func NewGithubService() (*GithubService, error) {
	tokenDb, errTok := db.InitTokenDb()
	githubDb, errGit := db.InitGithubDb()
	if err := cmp.Or(errTok, errGit); err != nil {
		return nil, err
	}
	return &GithubService{tokenDb: tokenDb, GithubDb: githubDb, reactService: nil}, nil
}

func (git *GithubService) InitReactClient(conn *grpc.ClientConn) {
	git.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (git *GithubService) storeNewWebHook(
	tokenInfo *models.Token,
	req *gRPCService.GitWebHookInfo,
	repoAction models.GAction,
	repoType models.GType,
	hookId int32,
) error {
	_, err := git.GithubDb.StoreNewGithub(&models.Github{
		ActionID:   uint(req.ActionId),
		UserID:     uint(tokenInfo.UserID),
		Activated:  true,
		RepoOwner:  req.Owner,
		RepoName:   req.Repo,
		RepoAction: repoAction,
		ActionType: repoType,
	})
	return err
}

func (git *GithubService) formatWebhookCallbackURL(event string, actionID uint32) (string, error) {
	envWebhookUrl, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(envWebhookUrl, "github", "push", actionID), nil
}
