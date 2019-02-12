package main

import (
	"time"

	"math/rand"

	"git.ucloudadmin.com/udb/v2/common"
	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud"
	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud/udemo"
	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
)

func EchoGray(reqBodyItf interface{}) interface{} {
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
		Name: proto.String(reqBody.GetName() + " (gray)"),
		Rc:   Rc,
	}
	return respBody
}

func init() {
	// register handler info
	HandlerInfoMap[common.GenerateGrayMessageType(udemo.MessageType_value["MY_ECHO_REQUEST"], 255)] = HandlerInfo{
		F:          EchoGray,
		ResponseID: udemo.MessageType_value["MY_ECHO_RESPONSE"],
		Timeout:    5 * time.Second,
	}
}