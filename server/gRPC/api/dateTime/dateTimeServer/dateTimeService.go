//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTimeService
//

package dateTime_server

import (
	"area/db"
	Dt_Types "area/gRPC/api/dateTime/dateTimeTypes"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	timeNowUrl = "https://tools.aimylogic.com/api/now?tz=Europe/Paris"
)

type DateTimeService struct {
	dtDb         *db.DateTimeDB
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
	dt := &DateTimeService{dtDb: DtDb, c: scheduler, reactService: nil}
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

	var dateData Dt_Types.AimylogicDateTime
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
	allDTActions, err := dt.dtDb.GetAllDTActionsActivated()

	if err != nil {
		log.Println(err)
		return
	}
	for _, dtAct := range *allDTActions {
		if dateData.Day == int(dtAct.DayMonth) && dateData.Hour == int(dtAct.Hours) && dateData.Minute == int(dtAct.Minutes) {
			ctx := grpcutils.CreateContextFromUserID(int(dtAct.UserID))
			dt.reactService.LaunchReaction(
				ctx,
				&gRPCService.LaunchRequest{ActionId: int64(dtAct.ActionID), PrevOutput: bytesBody},
			)
		}
	}
}

func (dt *DateTimeService) LaunchCronJob(ctx context.Context, req *gRPCService.TriggerTimeRequest) (*gRPCService.TriggerTimeResponse, error) {
	userID, errClaim := grpcutils.GetUserIdFromContext(ctx, "DateTimeService")
	if errClaim != nil {
		log.Println(ctx)
		return nil, errClaim
	}

	t, err := time.Parse(time.RFC3339, req.DatetimeString)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid format for date format needed: RFC3339")
	}
	log.Println(t)
	_, err = dt.dtDb.InsertNewDTAction(&models.DateTime{
		ActionID:  uint(req.ActionId),
		UserID:    userID,
		Activated: true,
		Minutes:   int32(t.Minute()),
		Hours:     int32(t.Hour()),
		DayMonth:  int64(t.Day()),
		Month:     int64(t.Month()),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &gRPCService.TriggerTimeResponse{}, nil
}

func (dt *DateTimeService) SetActivateAction(ctx context.Context, req *gRPCService.SetActivateTime) (*gRPCService.SetActivateTime, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "dt")
	if err != nil {
		return nil, err
	}
	log.Printf("Setting Action (%v) activated to %v\n", req.ActionId, req.Activated)
	data, err := dt.dtDb.SetActivateByActionID(req.Activated, userID, uint(req.ActionId))
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Could not set activated for time action")
	}
	log.Printf("Time Action with action ID: %v has activated state: %v", data.ActionID, data.Activated)
	return req, nil
}
