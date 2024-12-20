//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleService
//

package google

import (
	gRPCService "area/protogen/gRPC/proto"
	"context"
)

type GoogleService struct {
	// reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGoogleServiceServer
}

func (google *GoogleService) TestCall(context.Context, *gRPCService.TestRequest) (*gRPCService.TestRequest, error) {
	return nil, nil
}
