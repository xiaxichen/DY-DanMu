package rpcsupport

import (
	Log "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//ServeRpc:注册服务到rpc服务器
func ServeRpc(host string, services ...interface{}) error {
	for _, service := range services {
		//注册结构
		err := rpc.Register(service)
		if err != nil {
			return err
		}
	}
	Log.Infof("Run server on %s", host)
	// 创建一个新的监听
	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			Log.Errorf("accept error: %v", err)
			continue
		} else {
			Log.Infof("listening:%s", host)
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClinet(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
