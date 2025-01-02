//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleDriveComments
//

package google_server

import (
	"area/gRPC/api/google/drive"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (google *GoogleService) CreateCommentOnFile(ctx context.Context, req *gRPCService.CreateCommentReq) (*gRPCService.CreateCommentReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	listFiles, err := drive.ListFiles(tokenInfo.AccessToken)
	if err != nil {
		return nil, err
	}
	for _, f := range listFiles.Files {
		if f.Name == req.FileName {
			_, err = drive.CreateCommentOnFile(tokenInfo.AccessToken, f.ID, "content", drive.DriveComment{
				Content: req.Content,
			})
			return req, err
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "No such file: %v", req.FileName)
}

func deleteCorrectComment(tokenInfo *models.Token, f *drive.DriveFile, req *gRPCService.DeleteCommentReq) error {
	comments, err := drive.ListComments(tokenInfo.AccessToken, f.ID, "comments")
	if err != nil {
		log.Println("Error on list comments")
		return err
	}
	for _, com := range comments.Comments {
		com, err := drive.GetComment(tokenInfo.AccessToken, f.ID, com.ID, "content,id")
		if err != nil {
			continue
		}
		if com.Content == req.Content {
			return drive.DeleteCommentOnFile(tokenInfo.AccessToken, f.ID, com.ID, "")
		}
	}
	return status.Errorf(codes.InvalidArgument, "No such comment with content: %v", req.Content)
}

func (google *GoogleService) DeleteCommentOnFile(ctx context.Context, req *gRPCService.DeleteCommentReq) (*gRPCService.DeleteCommentReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	listFiles, err := drive.ListFiles(tokenInfo.AccessToken)
	if err != nil {
		log.Println("List files error")
		return nil, err
	}
	for _, f := range listFiles.Files {
		if f.Name == req.FileName {
			err := deleteCorrectComment(tokenInfo, &f, req)
			if err != nil {
				return nil, err
			}
			return req, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "No such file: %v", req.FileName)
}

func updateCorrectComment(tokenInfo *models.Token, f *drive.DriveFile, req *gRPCService.UpdateCommentReq) (*drive.DriveComment, error) {
	comments, err := drive.ListComments(tokenInfo.AccessToken, f.ID, "comments")
	if err != nil {
		log.Println("Error on list comments")
		return nil, err
	}
	for _, com := range comments.Comments {
		com, err := drive.GetComment(tokenInfo.AccessToken, f.ID, com.ID, "content,id")
		if err != nil {
			continue
		}
		log.Println(com.Content, req.OldContent)
		if com.Content == req.OldContent {
			return drive.UpdateCommentOnFile(tokenInfo.AccessToken, f.ID, com.ID, "content", drive.DriveComment{
				Content: req.NewContent,
			})
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "No such comment with content: %v", req.NewContent)
}

func (google *GoogleService) UpdateCommentOnFile(ctx context.Context, req *gRPCService.UpdateCommentReq) (*gRPCService.UpdateCommentReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if err != nil {
		return nil, err
	}
	listFiles, err := drive.ListFiles(tokenInfo.AccessToken)
	if err != nil {
		log.Println("List files error")
		return nil, err
	}
	for _, f := range listFiles.Files {
		if f.Name == req.FileName {
			_, err := updateCorrectComment(tokenInfo, &f, req)
			if err != nil {
				return nil, err
			}
			return req, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "No such file: %v", req.FileName)
}
