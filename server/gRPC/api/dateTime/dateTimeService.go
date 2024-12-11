//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTimeService
//

package dateTime

import (
	"area/db"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

const (
	timeNowUrl = "https://tools.aimylogic.com/api/now?tz=Europe/Paris"
)

type DateTimeService struct {
	db           *db.DateTimeDB
	c            *cron.Cron
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedDateTimeServiceServer
}

func NewDateTimeService() (*DateTimeService, error) {
	scheduler := cron.New()
	scheduler.Start()
	DtDb, err := db.InitDateTimeDb()

	if err != nil {
		return nil, err
	}
	dt := &DateTimeService{db: DtDb, c: scheduler, reactService: nil}
	dt.c.AddFunc("* * * * *", dt.checkTimeTrigger)
	return dt, nil
}

func (dt *DateTimeService) InitReactClient(conn *grpc.ClientConn) {
	dt.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (dt *DateTimeService) checkTimeTrigger() {
	r, err := http.Get(timeNowUrl)
	if err != nil {
		log.Println(err)
		return
	}
	bytesBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var dateData AimylogicDateTime
	err = json.Unmarshal(bytesBody, &dateData)
	if err != nil {
		log.Println(err)
		return
	}
	bytesBody, err = json.Marshal(&dateData)
	if err != nil {
		log.Println(err)
		return
	}
	allDTActions, err := dt.db.GetAllDTActionsActivated()

	if err != nil {
		log.Println(err)
		return
	}
	for _, dtAct := range *allDTActions {
		if dateData.Day == int(dtAct.DayMonth) && dateData.Hour == int(dtAct.Hours) && dateData.Minute == int(dtAct.Minutes) {
			dt.reactService.LaunchReaction(
				context.Background(),
				&gRPCService.LaunchRequest{ActionId: int64(dtAct.ActionID), PrevOutput: bytesBody},
			)
		}
	}
}

func (dt *DateTimeService) LaunchCronJob(_ context.Context, req *gRPCService.TriggerTimeRequest) (*gRPCService.TriggerTimeResponse, error) {
	log.Println("Starting cron job")
	dt.db.InsertNewDTAction(&models.DateTime{
		ActionID:  uint(req.ActionId),
		Activated: true,
		Minutes:   req.Minutes,
		Hours:     req.Hours,
		DayMonth:  req.DayMonth,
		Month:     req.Month,
		DayWeek:   req.DayWeek,
	})
	return &gRPCService.TriggerTimeResponse{}, nil
}
