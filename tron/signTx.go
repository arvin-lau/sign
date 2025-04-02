package tron

import (
	"crypto/sha256"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

func SignTransactionByKmService(transaction *core.Transaction, signType string, orderId string, coinType uint32, addressIdx uint32, wid uint32) (*core.Transaction, error) {
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, fmt.Errorf("proto marshal tx raw data error: %v", err)
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	switch signType {

	}
	kmCli := client2.GetKmClient()
	signature, err := kmCli.SignHash(walletId, types2.CoinTypeTRON, addrIndex, hash)
	if err != nil {
		log.Error().Msg("ERROR: tron SignTx kmCli.SignHash err:" + err.Error())
		return nil, err
	}
	transaction.Signature = append(transaction.Signature, signature)
	return transaction, nil
}

func SignTransactionByKmService1(transaction *core.Transaction, walletId uint32, addrIndex uint32) (*core.Transaction, error) {
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, fmt.Errorf("proto marshal tx raw data error: %v", err)
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	kmCli := client2.GetKmClient()
	signature, err := kmCli.SignHash(walletId, types2.CoinTypeTRON, addrIndex, hash)
	if err != nil {
		log.Error().Msg("ERROR: tron SignTx kmCli.SignHash err:" + err.Error())
		return nil, err
	}
	transaction.Signature = append(transaction.Signature, signature)
	return transaction, nil
}
