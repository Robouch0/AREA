//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// cryptoCheckTrigger
//

package cryptocompareserver

import (
	"area/models"
	service "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"log"
)

func (crypto *CryptoService) checkIsHigher() {
	actions, err := crypto.cryptoDb.GetActionsByType(models.IsHigher)
	if err != nil {
		log.Println("Error while getting actions")
		return
	}
	for _, act := range *actions {
		resp, err := GetCryptoInfos(&CryptoConfig{
			CryptoCurrency: act.CryptoCurrency,
			Currency: act.Currency,
		})
		if err != nil {
			log.Println(err)
			continue
		}
		threshold := resp.Currency[act.Currency]
		if uint(threshold) > act.Threshold {
			var cryptoP CryptoPayload
			cryptoP.CryptoCurrency = act.CryptoCurrency
			cryptoP.Currency = act.Currency
			cryptoP.Threshold = int(threshold)
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&cryptoP)
			if err != nil {
				log.Println("Could not marshal crypto current response")
				continue
			}
			crypto.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}

func (crypto *CryptoService) checkIsLower() {
	actions, err := crypto.cryptoDb.GetActionsByType(models.IsLower)
	if err != nil {
		log.Println("Error while getting actions")
		return
	}
	for _, act := range *actions {
		resp, err := GetCryptoInfos(&CryptoConfig{
			CryptoCurrency: act.CryptoCurrency,
			Currency: act.Currency,
		})
		if err != nil {
			log.Println(err)
			continue
		}
		threshold := resp.Currency[act.Currency]
		if uint(threshold) < act.Threshold {
			var cryptoP CryptoPayload
			cryptoP.CryptoCurrency = act.CryptoCurrency
			cryptoP.Currency = act.Currency
			cryptoP.Threshold = int(threshold)
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&cryptoP)
			if err != nil {
				log.Println("Could not marshal crypto current response")
				continue
			}
			crypto.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}

