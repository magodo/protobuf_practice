package main

import (
	"net"

	"fmt"

	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud"

	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
	ufnet "gitlab.ucloudadmin.com/udb/uframework/net"
	uftask "gitlab.ucloudadmin.com/udb/uframework/task"
)

func onDataIn(c *ufnet.TcpConnection, msg []byte) {
	fmt.Println("on data in")
	req := parseMsg(msg)
	fmt.Printf("receive message from %v:\n%v\n", c.Conn().RemoteAddr(), proto.MarshalTextString(req))
	message_type := req.GetHead().GetMessageType()
	task, err := uftask.NewTCPTask(int32(message_type))
	if err != nil {
		fmt.Printf("new task fail: %+v", err)
		return
	}

	res, err := task.Run(req)
	if err != nil {
		fmt.Printf("run task(%d) fail:%+v", task.Id, err)
		return
	}
	ufnet.SendTCPResponse(c, res)
}

func onDataOut(c net.Conn, msg []byte) {
	data := parseMsg(msg)
	fmt.Printf("Send message [%s => %s]:\n%v", c.LocalAddr(), c.RemoteAddr(), proto.MarshalTextString(data))
}

func parseMsg(msgRaw []byte) *ucloud.UMessage {
	msg := new(ucloud.UMessage)
	proto.Unmarshal(msgRaw, msg)
	return msg

}

func startNetworkService(listenIp string, listenPort int) error {
	ufnet.OnDataIn = onDataIn
	ufnet.OnDataOut = onDataOut
	return ufnet.ListenAndServeTCP(listenIp, listenPort)
}
