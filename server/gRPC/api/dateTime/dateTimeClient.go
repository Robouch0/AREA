//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTimeClient
//

package dateTime

import (
	gRPCService "area/protogen/gRPC/proto"

	"google.golang.org/grpc"
)

type DTServiceClient struct {
	gRPCService.DateTimeServiceClient
}

func NewDateTimeServiceClient(conn *grpc.ClientConn) *DTServiceClient {
	return &DTServiceClient{gRPCService.NewDateTimeServiceClient(conn)}
}

func (dt *DTServiceClient) SendAction(body map[string]any) (string, error) {

	
	// dt.LaunchCronJob(context.Background(), &gRPCService.TriggerTimeRequest{
	// Minutes:  1,
	// Hours:    -1,
	// DayMonth: -1,
	// Month:    -1,
	// DayWeek:  -1,
	// })
	return "", nil
}
