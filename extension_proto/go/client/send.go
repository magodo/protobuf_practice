package main

import (
	"log"
	"proto_foo/proto/foo"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
	ufnet "gitlab.ucloudadmin.com/udb/uframework/net"
)

type ValidateFunction func(resp *foo.Message) error

func DefaultValidater(resp *foo.Message) error {
	respID := resp.GetHeader().GetType()
	respBodyExt := MessageBodyExtensions[respID]
	respBodyRaw, err := proto.GetExtension(resp.GetBody(), respBodyExt)
	if err != nil {
		log.Fatal(err)
	}
	vResp := reflect.ValueOf(respBodyRaw).Elem()
	vRc := vResp.FieldByName("Rc").Elem()
	if *(vRc.FieldByName("RetCode").Interface().(*int32)) != 0 {
		return errors.New(*(vRc.FieldByName("RetMessage").Interface().(*string)))
	}
	return nil
}

var MessageBodyExtensions map[int32]*proto.ExtensionDesc

// if f == nil, it means no response is checked
func SendMessage(reqFullName string, reqBody interface{}, f ValidateFunction) error {

	reqID := getMsgID(reqFullName)
	req := &foo.Message{
		Header: &foo.Header{Type: proto.Int32(reqID)},
		Body:   &foo.Body{},
	}
	reqBodyExt := getMsgExtension(reqFullName)
	proto.SetExtension(req.GetBody(), reqBodyExt, reqBody)
	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	if f == nil {
		return ufnet.SendTCPRequestNoResponse(serverAddress.IP.String(), serverAddress.Port, reqRaw, 0)
	}

	respRaw, err := ufnet.SendTCPRequest(serverAddress.IP.String(), serverAddress.Port, reqRaw, 0)
	if err != nil {
		return err
	}

	resp := &foo.Message{}
	proto.Unmarshal(respRaw, resp)

	return f(resp)
}

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
	MessageBodyExtensions = proto.RegisteredExtensions((*foo.Body)(nil))
}
