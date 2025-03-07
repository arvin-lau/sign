package evm

import (
	"crypto/ecdsa"
	kmServiceCli "github.com/arvin-lau/sign/kmservice/client"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignTx(tx *types.Transaction, s types.Signer, prv *ecdsa.PrivateKey) (*types.Transaction, error) {
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	h := s.Hash(tx)
	kmCli := kmServiceCli.GetKmClient()
	//sig, err := kmCli.SignHash(wid, chain.GetCoinTypeByChainName(chainName), addressIdx, h.Bytes())
	sig, err := kmCli.SignHash(wid, kmserviceType.CoinType(coinType), addressIdx, h.Bytes())
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(s, sig)
}
