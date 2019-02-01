package main

import (
	"proto_foo/proto/foo"
	"proto_foo/proto/foo/ext"
	"time"

	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
)

func Foo1(reqBodyItf interface{}) interface{} {
	return &ext.FooResponse{
		Rc: &foo.Rc{
			RetCode:    proto.Int32(0),
			RetMessage: proto.String(""),
		},
	}
}

func init() {
	// register handler info
	HandlerInfoMap[ext.MessageType_value["FOO1_REQUEST"]] = HandlerInfo{
		F:          Foo1,
		ResponseID: ext.MessageType_value["FOO1_RESPONSE"],
		Timeout:    5 * time.Second,
	}
}
