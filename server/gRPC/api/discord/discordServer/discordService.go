//
// EPITECH PROJECT, 2024
// AREA
// File description:
// discordService
//

package discord_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DiscordService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedDiscordServiceServer
}

func NewDiscordService() (*DiscordService, error) {
	tokenDb, err := db.InitTokenDb()

	return &DiscordService{tokenDb: tokenDb, reactService: nil}, err
}

func (discord *DiscordService) CreateMessage(ctx context.Context, req *gRPCService.CreateMsg) (*gRPCService.CreateMsg, error) {
	if req.Channel == "" || req.Content == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	_, err := grpcutils.GetTokenByContext(ctx, discord.tokenDb, "DiscordService", "discord")
	if err != nil {
		return nil, err
	}

	BotToken, errToken := utils.GetEnvParameter("DISCORD_BOT_TOKEN")
	b, errMarshal := json.Marshal(req)
	if err = cmp.Or(errToken, errMarshal); err != nil {
		log.Println(err)
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages", req.Channel)
	putRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	putRequest.Header = http_utils.GetDefaultBotHTTPHeader(BotToken)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		io.Copy(os.Stderr, resp.Body)
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (discord *DiscordService) EditMessage(ctx context.Context, req *gRPCService.EditMsg) (*gRPCService.EditMsg, error) {
	if req.Channel == "" || req.MessageId == "" || req.Content == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	_, err := grpcutils.GetTokenByContext(ctx, discord.tokenDb, "DiscordService", "discord")
	if err != nil {
		return nil, err
	}

	BotToken, errToken := utils.GetEnvParameter("DISCORD_BOT_TOKEN")
	b, errMarshal := json.Marshal(req)
	if err = cmp.Or(errToken, errMarshal); err != nil {
		log.Println(err)
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v", req.Channel, req.MessageId)
	putRequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	putRequest.Header = http_utils.GetDefaultBotHTTPHeader(BotToken)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		io.Copy(os.Stderr, resp.Body)
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (discord *DiscordService) DeleteMessage(ctx context.Context, req *gRPCService.DeleteMsg) (*gRPCService.DeleteMsg, error) {
	if req.Channel == "" || req.MessageId == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	_, err := grpcutils.GetTokenByContext(ctx, discord.tokenDb, "DiscordService", "discord")
	if err != nil {
		return nil, err
	}

	BotToken, errToken := utils.GetEnvParameter("DISCORD_BOT_TOKEN")
	if err = cmp.Or(errToken); err != nil {
		log.Println(err)
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v", req.Channel, req.MessageId)
	putRequest, err := http.NewRequest("DELETE", url, nil)
	putRequest.Header = http_utils.GetDefaultBotHTTPHeader(BotToken)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 204 {
		io.Copy(os.Stderr, resp.Body)
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (discord *DiscordService) CreateReaction(ctx context.Context, req *gRPCService.CreateReact) (*gRPCService.CreateReact, error) {
	if req.Channel == "" || req.MessageId == "" || req.Emoji == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	_, err := grpcutils.GetTokenByContext(ctx, discord.tokenDb, "DiscordService", "discord")
	if err != nil {
		return nil, err
	}

	BotToken, errToken := utils.GetEnvParameter("DISCORD_BOT_TOKEN")

	if err = cmp.Or(errToken); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v/reactions/%v/@me", req.Channel, req.MessageId, req.Emoji)
	putRequest, err := http.NewRequest("PUT", url, nil)
	putRequest.Header = http_utils.GetDefaultBotHTTPHeader(BotToken)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 204 {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (discord *DiscordService) DeleteAllReactions(ctx context.Context, req *gRPCService.DeleteAllReact) (*gRPCService.DeleteAllReact, error) {
	if req.Channel == "" || req.MessageId == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	_, err := grpcutils.GetTokenByContext(ctx, discord.tokenDb, "DiscordService", "discord")
	if err != nil {
		return nil, err
	}

	BotToken, errToken := utils.GetEnvParameter("DISCORD_BOT_TOKEN")

	if err = cmp.Or(errToken); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v/reactions", req.Channel, req.MessageId)
	putRequest, err := http.NewRequest("DELETE", url, nil)
	putRequest.Header = http_utils.GetDefaultBotHTTPHeader(BotToken)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 204 {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (discord *DiscordService) DeleteReactions(ctx context.Context, req *gRPCService.DeleteReact) (*gRPCService.DeleteReact, error) {
	if req.Channel == "" || req.MessageId == "" || req.Emoji == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	_, err := grpcutils.GetTokenByContext(ctx, discord.tokenDb, "DiscordService", "discord")
	if err != nil {
		return nil, err
	}

	BotToken, errToken := utils.GetEnvParameter("DISCORD_BOT_TOKEN")

	if err = cmp.Or(errToken); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v/reactions/%v", req.Channel, req.MessageId, req.Emoji)
	putRequest, err := http.NewRequest("DELETE", url, nil)
	putRequest.Header = http_utils.GetDefaultBotHTTPHeader(BotToken)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 204 {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (discord *DiscordService) DeactivateAction(ctx context.Context, req *gRPCService.DeactivateDiscord) (*gRPCService.DeactivateDiscord, error) {
	return nil, status.Errorf(codes.Unavailable, "No Action for Discord Service yet")
}
