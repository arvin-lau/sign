package tron

import (
	"crypto/sha256"
	"fmt"
	"github.com/arvin-lau/sign/signature"
	kmserviceCli "github.com/arvin-lau/sign/signature/kmservice/client"
	kmserviceType "github.com/arvin-lau/sign/signature/kmservice/types"
	mpcCli "github.com/arvin-lau/sign/signature/mpcSign/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

func Sign(transaction *core.Transaction, signType string, orderId string, coinType uint32, addressIdx uint32, wid uint32) (*core.Transaction, error) {
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, fmt.Errorf("proto marshal tx raw data error: %v", err)
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	sig := make([]byte, 0)
	switch signType {
	case signature.SignByKmservice:
		kmCli := kmserviceCli.GetKmClient()
		sig, err = kmCli.SignHash(wid, kmserviceType.CoinType(coinType), addressIdx, hash)
		if err != nil {
			log.Error().Msg("ERROR: tron SignTx kmCli.SignHash err:" + err.Error())
			return nil, err
		}
	case signature.SignByMpc:
		kmCli := mpcCli.GetKmClient()
		sig, err = kmCli.SignHash(orderId, hash)
	}
	transaction.Signature = append(transaction.Signature, sig)
	return transaction, nil
}
