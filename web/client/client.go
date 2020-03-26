package client

import (
	"DY-DanMu/DMconfig/config"
	"DY-DanMu/persistServer/rpcsupport"
	"fmt"
	"net/rpc"
	"sync"
)

var lock sync.Mutex

//var wg = &sync.WaitGroup{}
var ClientRPC *rpc.Client

func init() {
	client, err := rpcsupport.NewClinet(config.DYWebConfig.Host)
	if err != nil {
		panic(err)
	}
	ClientRPC = client
}

func ReConnClientRPC(client *rpc.Client) (*rpc.Client, error) {
	lock.Lock()
	client, err := rpcsupport.NewClinet(config.DYWebConfig.Host)
	defer lock.Unlock()
	return client, err
}

// CheckError:处理连接断开异常
func CheckErrorForRPCDisconnect(err error) error {
	if fmt.Sprintf("%s", err) == "connection is shut down" {
		ClientRPC, err = ReConnClientRPC(ClientRPC)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}
