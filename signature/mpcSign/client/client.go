package client

import (
	"context"
	"errors"
	"github.com/arvin-lau/sign/signature/mpcSign/types"
	zlog "github.com/rs/zerolog/log"
	"github.com/smallnest/rpcx/client"
	"time"
)

type MpcSignClient struct {
	server  []string
	xclient client.XClient
}

// NewMpcSignClient 初始化服务器
func NewMpcSignClient(server []string, token string) *MpcSignClient {
	kbpair := []*client.KVPair{}
	for _, s := range server {
		kbpair = append(kbpair, &client.KVPair{
			Key: s,
		})
	}
	d, _ := client.NewMultipleServersDiscovery(kbpair)
	xclient := client.NewXClient("Arith", client.Failover, client.RandomSelect, d, client.DefaultOption)
	xclient.Auth(token)
	return &MpcSignClient{
		server:  server,
		xclient: xclient,
	}
}

func (k *MpcSignClient) Close() {
	err := k.xclient.Close()
	if err != nil {
		zlog.Error().Msg("err: " + err.Error())
	}
}

func (k *MpcSignClient) SignHash(orderId string, hash []byte) (sig []byte, err error) {
	args := types.SignReq{
		OrderId: orderId,
		Data:    hash,
	}
	reply := &types.SignRsp{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*180)
	err = k.xclient.Call(ctx, runFuncName(), args, reply)
	if err != nil {
		zlog.Error().Msg("failed to call: " + err.Error())
		return
	}
	if reply.Err != "" {
		return reply.Sign, errors.New(reply.Err)
	}
	return reply.Sign, nil
}
