//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTimeService
//

package dateTime

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

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
	// Store in database the data of the trigger request
	fmt.Println("Starting cron job")
	// dt.c.AddFunc("* * * * *",
	func() { // Format this correctly
		log.Println("Trigger activated")
		r, err := http.Get("https://tools.aimylogic.com/api/now?tz=Europe/Paris")
		if err != nil {
			log.Println(err)
			return
		}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		var dateData AimylogicDateTime
		err = json.Unmarshal(b, &dateData)
		if err != nil {
			log.Println(err)
			return
		}
		b, err = json.Marshal(&dateData)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("LaunchReaction !")
		dt.reactService.LaunchReaction(
			context.Background(),
			&gRPCService.LaunchRequest{ActionId: int64(req.ActionId), PrevOutput: b},
		)
	}()
	// )

	return &gRPCService.TriggerTimeResponse{}, nil
}
