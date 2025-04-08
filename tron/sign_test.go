package tron

import (
	"fmt"
	"github.com/arvin-lau/sign/signature"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	//https://www.trongrid.io/  login account
	nodeUrl := "your tronGrid node url"
	apiKey := "your tronGrid ap key"
	if err := NewTronCli(nodeUrl, apiKey); err != nil {
		panic(err)
	}
	tronCli := GetTronCli()
	txExtention, err := tronCli.Transfer("TP4WwcyEbXtvNu1igPiaz7LxHYPqdE3ggc", "TJRTf1t6xaiGuaxqUtnaUy6yEYVyFKABVm", 1)
	if err != nil {
		log.Error().Msgf("tronCli.Transfer err: %v", err)
		return
	}
	tx, err := Sign(txExtention.Transaction, signature.SignByKmservice, "", uint32(0x800000c3), 0, 1)
	if err != nil {
		log.Error().Msgf("Sign err: %v", err)
		return
	}
	resp, err := tronCli.Broadcast(tx)
	if err != nil {
		log.Error().Msgf("Broadcast err: %v", err)
		return
	}
	log.Info().Msgf("broadcast resp: %v", resp)

}

var cli *client.GrpcClient

func GetTronCli() *client.GrpcClient {
	return cli
}

func NewTronCli(nodeUrl, apiKey string) error {
	cli = client.NewGrpcClient(nodeUrl)
	cli.SetTimeout(time.Second * 30)
	if err := cli.SetAPIKey(apiKey); err != nil {
		return err
	}

	if err := cli.Start(grpc.WithInsecure()); err != nil {
		return fmt.Errorf("grpc client start error: %v", err)
	}
	return nil
}
