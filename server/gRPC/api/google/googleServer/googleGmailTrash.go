//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleGmailTrash
//

package google_server

import (
	"area/gRPC/api/google/gmail"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"

	"context"
	"fmt"
	"log"
	"slices"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func trashExpr(trash bool) string {
	if trash {
		return ""
	}
	return "TRASH"
}

func (google *GoogleService) moveTrashImplem(ctx context.Context, req *gRPCService.TrashEmailRequestMe, trash bool) (*gRPCService.TrashEmailRequestMe, error) {
	if req.Subject == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Empty subject for move to trash")
	}
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleGmailService", "google")
	if err != nil {
		return nil, err
	}

	emails, err := gmail.GetListEmails("me", tokenInfo.AccessToken, trashExpr(trash))
	if err != nil {
		return nil, err
	}

	for _, message := range emails.Messages {
		mess, err := gmail.GetEmail("me", tokenInfo.AccessToken, message.ID, "metadata", "subject")
		if err != nil {
			log.Println(err)
			continue
		}
		idx := slices.IndexFunc[[]gmail.GmailHeader](mess.Payload.Headers, func(h gmail.GmailHeader) bool {
			return h.Name == "Subject"
		})
		if idx == -1 {
			continue
		}
		if mess.Payload.Headers[idx].Value == req.Subject {
			m, err := gmail.MoveEmail(tokenInfo.AccessToken, "me", message.ID, trash)
			if err != nil {
				return nil, err
			}
			log.Println(m.ID)
			return req, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Did not find any email with subject %s", req.Subject))
}

func (google *GoogleService) MoveToTrash(ctx context.Context, req *gRPCService.TrashEmailRequestMe) (*gRPCService.TrashEmailRequestMe, error) {
	return google.moveTrashImplem(ctx, req, true)
}

func (google *GoogleService) MoveFromTrash(ctx context.Context, req *gRPCService.TrashEmailRequestMe) (*gRPCService.TrashEmailRequestMe, error) {
	return google.moveTrashImplem(ctx, req, false)
}
