package main

import (
	"proto_foo/proto/foo/ext"

	"time"

	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
)

func Echo(reqBodyItf interface{}) interface{} {
	reqBody := reqBodyItf.(*ext.EchoRequest)
	respBody := &ext.EchoResponse{
		Msg: proto.String(reqBody.GetMsg()),
	}
	return respBody
}

func init() {
	// register handler info
	HandlerInfoMap[ext.MessageType_value["ECHO_REQUEST"]] = HandlerInfo{
		F:          Echo,
		ResponseID: ext.MessageType_value["ECHO_RESPONSE"],
		Timeout:    5 * time.Second,
	}
}
