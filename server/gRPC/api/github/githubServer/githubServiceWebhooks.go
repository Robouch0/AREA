//
// EPITECH PROJECT, 2025
// AREA
// File description:
// githubServiceWebhooks
//

package github

import (
	github_webhooks "area/gRPC/api/github/githubWebhooks"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	webhookURL = "https://api.github.com/repos/%v/%v/hooks"
)

func (git *GithubService) CreatePushWebhook(ctx context.Context, req *gRPCService.GitWebHookInfo) (*gRPCService.GitWebHookInfo, error) {
	if req.Owner == "" || req.Repo == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument for webhook repo")
	}
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GithubService", "github")
	if err != nil {
		return nil, err
	}
	formattedWebhookURL, err := git.formatWebhookCallbackURL("push", uint32(req.ActionId))
	if err != nil {
		return nil, err
	}

	err = github_webhooks.SendCreateWebHook(tokenInfo, req.Owner, req.Repo, webhookURL, &github_webhooks.GitWebHookRequest{
		Event:  []string{"push"},
		Config: github_webhooks.GithubConfig{Url: formattedWebhookURL, Content: "json"},
	})
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, models.GCreate); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GithubService) CreateDeleteBranchWebhook(ctx context.Context, req *gRPCService.GitWebHookInfo) (*gRPCService.GitWebHookInfo, error) {
	if req.Owner == "" || req.Repo == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument for webhook repo")
	}
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GithubService", "github")
	if err != nil {
		return nil, err
	}
	formattedWebhookURL, err := git.formatWebhookCallbackURL("push", uint32(req.ActionId))
	if err != nil {
		return nil, err
	}

	err = github_webhooks.SendCreateWebHook(tokenInfo, req.Owner, req.Repo, webhookURL, &github_webhooks.GitWebHookRequest{
		Event:  []string{"delete"},
		Config: github_webhooks.GithubConfig{Url: formattedWebhookURL, Content: "json"},
	})
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, models.GDelete); err != nil {
		return nil, err
	}
	return req, nil
}

func (git *GithubService) CreateForkRepositoryWebhook(ctx context.Context, req *gRPCService.GitWebHookInfo) (*gRPCService.GitWebHookInfo, error) {
	if req.Owner == "" || req.Repo == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument for webhook repo")
	}
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, git.tokenDb, "GithubService", "github")
	if err != nil {
		return nil, err
	}
	formattedWebhookURL, err := git.formatWebhookCallbackURL("fork", uint32(req.ActionId))
	if err != nil {
		return nil, err
	}

	err = github_webhooks.SendCreateWebHook(tokenInfo, req.Owner, req.Repo, webhookURL, &github_webhooks.GitWebHookRequest{
		Event:  []string{"fork"},
		Config: github_webhooks.GithubConfig{Url: formattedWebhookURL, Content: "json"},
	})
	if err != nil {
		return nil, err
	}
	if err := git.storeNewWebHook(tokenInfo, req, models.GFork); err != nil {
		return nil, err
	}
	return req, nil
}

/// Deactivate Webhook

func (github *GithubService) SetActivateAction(ctx context.Context, req *gRPCService.SetActivateGithub) (*gRPCService.SetActivateGithub, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, github.tokenDb, "GithubService", "github")
	if err != nil {
		return nil, err
	}
	action, err := github.GithubDb.GetGithubByActionID(uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	if !req.Activated {
		err = github_webhooks.UpdatePushWebhook(tokenInfo, ctx, action, false)
	} else {
		err = github_webhooks.UpdatePushWebhook(tokenInfo, ctx, action, true)
	}
	if err != nil {
		return nil, err
	}
	_, err = github.GithubDb.SetActivateByActionID(req.Activated, uint(tokenInfo.UserID), uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}
