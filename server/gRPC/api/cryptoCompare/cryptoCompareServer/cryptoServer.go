//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// cryptoServer
//

package cryptocompareserver

import (
	"area/db"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"cmp"
	"context"
	"log"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CryptoService struct {
	tokenDb      *db.TokenDb
	cryptoDb     *db.CryptoCompareDB
	c            *cron.Cron
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedCryptoServiceServer
}

func NewCryptoService() (*CryptoService, error) {
	scheduler := cron.New()
	scheduler.Start()
	tokenDb, errTok := db.InitTokenDb()
	cryptoDb, errW := db.InitCryptoCompareDb()
	if err := cmp.Or(errTok, errW); err != nil {
		return nil, err
	}
	CryptoCompare := &CryptoService{
		tokenDb:      tokenDb,
		cryptoDb:     cryptoDb,
		c:            scheduler,
		reactService: nil,
	}
	CryptoCompare.c.AddFunc("* * * * *", CryptoCompare.checkIsHigher)
	CryptoCompare.c.AddFunc("* * * * *", CryptoCompare.checkIsLower)
	return CryptoCompare, nil
}

func (crypto *CryptoService) InitReactClient(conn *grpc.ClientConn) {
	crypto.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (crypto *CryptoService) createNewCryptoInfo(
	userID uint,
	actionID int,
	actionType models.CryptoActionType,
	cryptocurr string,
	currency string,
	threshold float64,
) error {
	_, err := crypto.cryptoDb.InsertNewCryptoCompare(&models.CryptoCompare{
		ActionID:       uint(actionID),
		UserID:         userID,
		ActionType:     actionType,
		Activated:      true,
		CryptoCurrency: cryptocurr,
		Currency:       currency,
		Threshold:      uint(threshold),
	})
	return err
}

func (crypto *CryptoService) IsHigherThanTrigger(ctx context.Context, req *gRPCService.IsHigherThanTriggerReq) (*gRPCService.IsHigherThanTriggerReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "CryptoService")
	if err != nil {
		return nil, err
	}
	if req.CryptoCurrency == "" || req.Currency == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing argument for crypto service")
	}

	resp, err := GetCryptoInfos(&CryptoConfig{
		CryptoCurrency: req.CryptoCurrency,
		Currency:       req.Currency,
	})
	if err != nil {
		log.Println("Could not fetch weather data: ", err)
		return nil, err
	}
	t := resp.Currency[req.Currency]
	err = crypto.createNewCryptoInfo(userID, int(req.ActionId), models.IsHigher, req.CryptoCurrency, req.Currency, t)
	if err != nil {
		return nil, err
	}
	log.Printf("Crypto currency %v is being looked at\n", req.CryptoCurrency)
	return req, nil
}

func (crypto *CryptoService) IsLowerThanTriggerReq(ctx context.Context, req *gRPCService.IsLowerThanTriggerReq) (*gRPCService.IsLowerThanTriggerReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "CryptoService")
	if err != nil {
		return nil, err
	}
	if req.CryptoCurrency == "" || req.Currency == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing argument for crypto service")
	}

	resp, err := GetCryptoInfos(&CryptoConfig{
		CryptoCurrency: req.CryptoCurrency,
		Currency:       req.Currency,
	})
	if err != nil {
		log.Println("Could not fetch weather data: ", err)
		return nil, err
	}
	t := resp.Currency[req.Currency]
	err = crypto.createNewCryptoInfo(userID, int(req.ActionId), models.IsHigher, req.CryptoCurrency, req.Currency, t)
	if err != nil {
		return nil, err
	}
	log.Printf("Crypto currency %v is being looked at\n", req.CryptoCurrency)
	return req, nil
}

func (crypto *CryptoService) SetActivate(ctx context.Context, req *gRPCService.SetActivateCrypto) (*gRPCService.SetActivateCrypto, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "crypto")
	if err != nil {
		return nil, err
	}
	_, err = crypto.cryptoDb.SetActivateByActionID(req.Activated, userID, uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (crypto *CryptoService) DeleteAction(ctx context.Context, req *gRPCService.DeleteActionReq) (*gRPCService.DeleteActionReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "crypto")
	if err != nil {
		return nil, err
	}
	return req, crypto.cryptoDb.DeleteByActionID(userID, uint(req.ActionId))
}
