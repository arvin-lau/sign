package evm

import (
	"context"
	"fmt"
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
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := SignTx(tx, types.NewEIP155Signer(chainID), SignByKmservice, "", uint32(0x8000003c), 1, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
