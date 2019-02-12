package main

import (
	"time"

	"math/rand"

	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud"
	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud/udemo"
	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
)

func Echo(reqBodyItf interface{}) interface{} {
	reqBody := reqBodyItf.(*udemo.MyEchoRequest)

	Rc := &ucloud.ResponseCode{}
	if rand.Intn(2)%2 == 0 {
		Rc.Retcode = proto.Int32(0)
	} else {
		Rc.Retcode = proto.Int32(-1)
		Rc.ErrorMessage = proto.String("bad luck...")
	}

	respBody := &udemo.MyEchoResponse{
		Id:   proto.String(reqBody.GetId()),
		Name: proto.String(reqBody.GetName()),
		Rc:   Rc,
	}
	return respBody
}

func init() {
	// register handler info
	HandlerInfoMap[udemo.MessageType_value["MY_ECHO_REQUEST"]] = HandlerInfo{
		F:          Echo,
		ResponseID: udemo.MessageType_value["MY_ECHO_RESPONSE"],
		Timeout:    5 * time.Second,
	}

	rand.Seed(0)
}
