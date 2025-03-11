package client

import (
	"context"
	"errors"
	"github.com/arvin-lau/sign/kmservice/types"
	zlog "github.com/rs/zerolog/log"
	"github.com/smallnest/rpcx/client"
)

type KmClient struct {
	server  []string
	xclient client.XClient
}

// NewKmClient 初始化服务器
func NewKmClient(server []string, token string) *KmClient {
	kbpair := []*client.KVPair{}
	for _, s := range server {
		kbpair = append(kbpair, &client.KVPair{
			Key: s,
		})
	}
	d, _ := client.NewMultipleServersDiscovery(kbpair)
	xclient := client.NewXClient("Arith", client.Failover, client.RandomSelect, d, client.DefaultOption)
	xclient.Auth(token)
	return &KmClient{
		server:  server,
		xclient: xclient,
	}
}

func (k *KmClient) Close() {
	err := k.xclient.Close()
	if err != nil {
		zlog.Error().Msg("err" + err.Error())
	}
}
func (k *KmClient) GetAccountAddress(wid uint32, coin types.CoinType, addressidx uint32) (string, string, error) {
	args := types.ArgsGetAddress{
		Wid:          wid,
		CoinType:     coin,
		AddressIndex: addressidx,
	}
	reply := &types.ReplyGetAddress{}
	err := k.xclient.Call(context.Background(), runFuncName(), args, reply)
	address := reply.Address
	pK := reply.Pubkey
	return address, pK, err
}

func (k *KmClient) SignHash(wid uint32, coin types.CoinType, addressidx uint32, hash []byte, lenVin ...int) (sig []byte, err error) {
	args := types.ArgsSign{
		Wid:          wid,
		CoinType:     coin,
		AddressIndex: addressidx,
		Hash:         hash,
	}
	if len(lenVin) != 0 {
		args.LenVin = lenVin[0]
	}
	reply := &types.ReplySign{}
	err = k.xclient.Call(context.Background(), runFuncName(), args, reply)
	if err != nil {
		zlog.Error().Msg("failed to call: %v" + err.Error())
		return
	}
	if reply.Err != "" {
		return reply.Result, errors.New(reply.Err)
	}
	return reply.Result, nil
}
