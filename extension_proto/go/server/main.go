package main

import (
	"log"

	uftask "gitlab.ucloudadmin.com/udb/uframework/task"
)

func main() {

	// 注册服务
	for reqID, _ := range HandlerInfoMap {
		uftask.RegisterTCPTaskHandle(reqID, uftask.TCPTaskFunc(HandleRequest), HandlerInfoMap[reqID].Timeout)
	}

	// 启动TCP服务
	if err := startNetworkService("127.0.0.1", 8888); err != nil {
		log.Fatalf("%+v", err)
	}
}
