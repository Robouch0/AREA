//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleGmailLabels
//

package google_server

import (
	"area/gRPC/api/google/gmail"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
)

func (google *GoogleService) CreateLabel(ctx context.Context, req *gRPCService.CreateLabelReq) (*gRPCService.CreateLabelReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	_, err = gmail.CreateLabel(tokenInfo.AccessToken, "me", gmail.GmailLabel{
		Name:                  req.Name,
		MessageListVisibility: gmail.GmailMessageVisibility(req.MessageListVisibility),
		LabelListVisibility:   gmail.GmailLabelVisibility(req.LabelListVisibility),
		Type:                  gmail.GmailLabelType(req.Type),
	})
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (google *GoogleService) UpdateLabel(ctx context.Context, req *gRPCService.UpdateLabelReq) (*gRPCService.UpdateLabelReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	_, err = gmail.PutLabel(tokenInfo.AccessToken, "me", req.OldName, gmail.GmailLabel{
		Name:                  req.NewName,
		MessageListVisibility: gmail.GmailMessageVisibility(req.MessageListVisibility),
		LabelListVisibility:   gmail.GmailLabelVisibility(req.LabelListVisibility),
		Type:                  gmail.GmailLabelType(req.Type),
	})
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (google *GoogleService) DeleteLabel(ctx context.Context, req *gRPCService.DeleteLabelReq) (*gRPCService.DeleteLabelReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	if err := gmail.DeleteLabel(tokenInfo.AccessToken, "me", req.Name); err != nil {
		return nil, err
	}
	return req, nil
}
