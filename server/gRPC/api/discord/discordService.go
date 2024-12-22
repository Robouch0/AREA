//
// EPITECH PROJECT, 2024
// AREA
// File description:
// discordService
//

package discord

import (
	// "area/db"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type DiscordService struct {
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedDiscordServiceServer
}

func NewDiscordService() DiscordService {
	return DiscordService{reactService: nil}
}

func (disCli *DiscordService) CreateMessage(_ context.Context, req *gRPCService.CreateMsg) (*gRPCService.CreateMsg, error) {
	if req.Channel == "" || req.Content == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	// Tokendb := db.InitTokenDb()
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages", req.Channel)
	putRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	putRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (disCli *DiscordService) EditMessage(_ context.Context, req *gRPCService.EditMsg) (*gRPCService.EditMsg, error) {
	if req.Channel == "" || req.MessageId == "" || req.Content == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	// Tokendb := db.InitTokenDb()
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v", req.Channel, req.MessageId)
	putRequest, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	putRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (disCli *DiscordService) DeleteMessage(_ context.Context, req *gRPCService.DeleteMsg) (*gRPCService.DeleteMsg, error) {
	if req.Channel == "" || req.MessageId == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	// Tokendb := db.InitTokenDb()
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v", req.Channel, req.MessageId)
	putRequest, err := http.NewRequest("DELETE", url, bytes.NewBuffer(b))
	putRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (disCli *DiscordService) CreateReaction(_ context.Context, req *gRPCService.CreateReact) (*gRPCService.CreateReact, error) {
	if req.Channel == "" || req.MessageId == "" || req.Emoji == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	// Tokendb := db.InitTokenDb()
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v/reactions/%v/@me", req.Channel, req.MessageId, req.Emoji)
	putRequest, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	putRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "204 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (disCli *DiscordService) DeleteAllReactions(_ context.Context, req *gRPCService.DeleteAllReact) (*gRPCService.DeleteAllReact, error) {
	if req.Channel == "" || req.MessageId == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	// Tokendb := db.InitTokenDb()
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v/reactions/", req.Channel, req.MessageId)
	putRequest, err := http.NewRequest("DELETE", url, bytes.NewBuffer(b))
	putRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "204 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

func (disCli *DiscordService) DeleteReactions(_ context.Context, req *gRPCService.DeleteReact) (*gRPCService.DeleteReact, error) {
	if req.Channel == "" || req.MessageId == "" || req.Emoji == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	// Tokendb := db.InitTokenDb()
	bearerTok, err := utils.GetEnvParameterToBearer("API_GITHUB")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://discord.com/api/channels/%v/messages/%v/reactions/%v", req.Channel, req.MessageId, req.Emoji)
	putRequest, err := http.NewRequest("DELETE", url, bytes.NewBuffer(b))
	putRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
	cli := &http.Client{}
	resp, err := cli.Do(putRequest)
	if err != nil {
		return nil, err
	}
	if resp.Status != "204 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}

/*
id
channel_id
author [1]
content [2]
timestamp
edited_timestamp
tts
mention_everyone
mentions
mention_roles
attachments [2]
embeds [2]
pinned
type
*/