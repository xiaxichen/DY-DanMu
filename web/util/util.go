package util

import "DY-DanMu/web/client"

// RpcClientShutDownErrorhandler:处理连接被shutdonwn的异常
func RpcClientShutDownErrorhandler(err error) error {
	if err != nil {
		err = client.CheckErrorForRPCDisconnect(err)
		return nil
	}
	return err
}
