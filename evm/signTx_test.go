package evm

import (
	"context"
	"fmt"
	"github.com/arvin-lau/sign/signature"
	kmservice "github.com/arvin-lau/sign/signature/kmservice/client"
	mpc "github.com/arvin-lau/sign/signature/mpcSign/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
)

func TestSignTx(t *testing.T) {
	kmservice.InitKmClient([]string{"192.168.50.89:9093"}, "token")
	mpc.InitMpcSignClient([]string{"192.168.50.89:9093"}, "token")
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress("0xa83114A443dA1CecEFC50368531cACE9F37fCCcb")
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	var data []byte
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	//eip 1559
	//inner := &types.DynamicFeeTx{
	//	ChainID:   chainID,
	//	Nonce:     nonce,
	//	GasTipCap: gasPrice,
	//	GasFeeCap: gasFeeCap,
	//	Gas:       gasLimit,
	//	To:        &toAddress,
	//	Value:     value,
	//	Data:      data,
	//}
	//tx := types.NewTx(inner)

	signedTx, err := SignTx(tx, types.NewEIP155Signer(chainID), signature.SignByKmservice, "", uint32(0x8000003c), 1, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

//func a() {
//	inner := &types.DynamicFeeTx{
//		ChainID:   chainId,
//		Nonce:     nonce,
//		GasTipCap: gasPrice,
//		GasFeeCap: gasFeeCap,
//		Gas:       200000000,
//		To:        &toAddress,
//		Value:     big.NewInt(1),
//	}
//	fromAddress, privateKey, _ := utils.GetAddress(fromAddrIndex)
//	fmt.Println("fromAddress: ", fromAddress)
//	unsignTx := types.NewTx(inner)
//}
