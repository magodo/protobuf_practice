package main

import (
	"strings"

	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud"

	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
)

var MessageBodyExtensions map[int32]*proto.ExtensionDesc

func getMsgID(msgFullName string) int32 {
	msgTypeName := getMsgDomain(msgFullName) + ".MessageType"
	msgName := getMsgName(msgFullName)
	return proto.EnumValueMap(msgTypeName)[strings.ToUpper(msgName)]
}

func getMsgDomain(msgFullName string) string {
	items := strings.Split(msgFullName, ".")
	return strings.Join(items[0:len(items)-1], ".")
}

func getMsgName(msgFullName string) string {
	items := strings.Split(msgFullName, ".")
	return items[len(items)-1]
}

func getMsgExtension(msgFullName string) (extension *proto.ExtensionDesc) {
	msgID := getMsgID(msgFullName)
	extension = MessageBodyExtensions[msgID]
	return
}

func init() {
	MessageBodyExtensions = proto.RegisteredExtensions((*ucloud.Body)(nil))
}
