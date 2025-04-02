package evm

import (
	"fmt"
	"github.com/arvin-lau/sign"
	kmserviceCli "github.com/arvin-lau/sign/signature/kmservice/client"
	kmserviceType "github.com/arvin-lau/sign/signature/kmservice/types"
	mpcCli "github.com/arvin-lau/sign/signature/mpcSign/client"
	"github.com/ethereum/go-ethereum/core/types"
)

func SignTx(tx *types.Transaction, s types.Signer, signType string, orderId string, coinType uint32, addressIdx uint32, wid uint32) (signedTx *types.Transaction, err error) {
	h := s.Hash(tx)
	sig := make([]byte, 0)
	switch signType {
	case sign.SignByKmservice:
		kmCli := kmserviceCli.GetKmClient()
		sig, err = kmCli.SignHash(wid, kmserviceType.CoinType(coinType), addressIdx, h.Bytes())
	case sign.SignByMpc:
		kmCli := mpcCli.GetKmClient()
		sig, err = kmCli.SignHash(orderId, h.Bytes())
	default:
		return nil, fmt.Errorf("%v signType not support", signType)
	}
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(s, sig)
}
