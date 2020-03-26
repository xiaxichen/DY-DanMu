package rpcsupport

import (
	"fmt"
	Log "github.com/sirupsen/logrus"
	"net/rpc"
)

func CreateClientPool(hosts []string) chan *rpc.Client {
	var rpcClient []*rpc.Client
	for _, host := range hosts {
		client, err := NewClinet(host)
		if err != nil {
			Log.Errorf(fmt.Sprintf("server client error! ipAddress:%s \n", host))
			continue
		}
		Log.Infof(fmt.Sprintf("client create for ipAddress:%s \n", host))
		rpcClient = append(rpcClient, client)
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range rpcClient {
				out <- client
			}
		}
	}()
	return out
}
