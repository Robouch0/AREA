//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTimeService
//

package api

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

// type msg struct {
// 	Msg string `json:"msg"`
// }

type DTServiceClient struct {
	gRPCService.DateTimeServiceClient
}

func NewDateTimeServiceClient(conn *grpc.ClientConn) *DTServiceClient {
	return &DTServiceClient{gRPCService.NewDateTimeServiceClient(conn)}
}

func (dt *DTServiceClient) SendAction(body []byte) (string, error) {
	// msg := new(msg)
	// err := json.Unmarshal([]byte(body), msg)
	//
	// r, err := hello.SayHello(context.Background(), &gRPCService.HelloWorldRequest{Message: msg.Msg})
	// if err != nil {
	// return "", err
	// }
	// return r.GetMessage(), nil
	dt.LaunchCronJob(context.Background(), &gRPCService.TriggerTimeRequest{
		Minutes:  1,
		Hours:    -1,
		DayMonth: -1,
		Month:    -1,
		DayWeek:  -1,
	})
	return "", nil
}

////

type DateTimeService struct {
	db           *db.UserDb
	c            *cron.Cron
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedDateTimeServiceServer
}

func NewDateTimeService(db *db.UserDb) DateTimeService {
	scheduler := cron.New()
	scheduler.Start()
	return DateTimeService{db: db, c: scheduler, reactService: nil}
}

func (dt *DateTimeService) InitReactClient(conn *grpc.ClientConn) {
	dt.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (dt *DateTimeService) LaunchCronJob(_ context.Context, req *gRPCService.TriggerTimeRequest) (*gRPCService.TriggerTimeResponse, error) {
	// Format this correctly
	dt.c.AddFunc("* * * * *", func() {
		log.Println("Trigger activated")
		dt.reactService.LaunchReaction(context.Background(), &gRPCService.ReactionRequest{Msg: "Hello"})
	})

	fmt.Println("Starting cron job")
	return &gRPCService.TriggerTimeResponse{}, nil
}
