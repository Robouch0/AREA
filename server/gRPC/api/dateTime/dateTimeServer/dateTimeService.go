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
	"strconv"
	"strings"

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
	bytesBody, err = json.Marshal(&dateData) // Check but not really usefull
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

    log.Println(req)
    parts := strings.Split(req.DatetimeString, "/")
    if len(parts) != 3 {
        log.Println(req.DatetimeString)
        log.Println("Error while parsing string dateTime, not enough arguments")
        return nil, nil
    }

    dayMonth, err := strconv.Atoi(parts[0])
    if err != nil {
        log.Println("Error while parsing month part")
        return nil, err
    }

    month, err := strconv.Atoi(parts[1])
    if err != nil {
        log.Println("Error while parsing day part")
        return nil, err
    }

    timeParts := strings.Split(parts[2], ":")
    if len(timeParts) != 2 {
        log.Println("Error while parsing hours part, no arguments")
        return nil, err
    }

    hours, err := strconv.Atoi(timeParts[0])
    if err != nil {
        log.Println("Error while parsing hours part")
        return nil, err
    }

    minutes, err := strconv.Atoi(timeParts[1])
    if err != nil {
        log.Println("Error while parsing minutes part")
        return nil, err
    }

    log.Println("Starting cron job")
    _, err = dt.dtDb.InsertNewDTAction(&models.DateTime{
        ActionID:  uint(req.ActionId),
        UserID:    userID,
        Activated: true,
        Minutes:   int32(minutes),
        Hours:     int32(hours),
        DayMonth:  int64(dayMonth),
        Month:     int64(month),
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
	resp, err := dt.dtDb.SetActivateByActionID(req.Activated, userID, uint(req.ActionId))
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Could not set activated for time action")
	}
	log.Printf("Time Action with action ID: %v has activated state: %v", resp.ActionID, resp.Activated)
	return req, nil
}
