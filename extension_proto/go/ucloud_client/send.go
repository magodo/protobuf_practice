package main

import (
	"log"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"

	"git.ucloudadmin.com/udb/v2/common"
	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud"

	_ "gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud/udemo"

	ufmessage "gitlab.ucloudadmin.com/udb/uframework/message"
	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
	ufnet "gitlab.ucloudadmin.com/udb/uframework/net"
)

type ValidateFunction func(resp *ucloud.UMessage) error

func DefaultValidater(resp *ucloud.UMessage) error {
	respID := resp.GetHead().GetMessageType()
	respBodyExt := MessageBodyExtensions[respID]
	respBodyRaw, err := proto.GetExtension(resp.GetBody(), respBodyExt)
	if err != nil {
		log.Fatal(err)
	}
	vResp := reflect.ValueOf(respBodyRaw).Elem()
	vRc := vResp.FieldByName("Rc").Elem()
	if *(vRc.FieldByName("Retcode").Interface().(*int32)) != 0 {
		return errors.New(*(vRc.FieldByName("ErrorMessage").Interface().(*string)))
	}
	return nil
}

var MessageBodyExtensions map[int32]*proto.ExtensionDesc

// if f == nil, it means no response is checked
func SendCheckUMessage(dest net.TCPAddr, realReqID int32, reqBody interface{}, uuid string, timeout time.Duration, f ValidateFunction) error {

	reqID, _ := common.GetMessageGrayInfo(realReqID)
	req := ufmessage.NewMessage(realReqID, uuid, false, 1, 0, "")
	reqBodyExt := MessageBodyExtensions[reqID]
	proto.SetExtension(req.GetBody(), reqBodyExt, reqBody)
	reqRaw, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	if f == nil {
		return ufnet.SendTCPRequestNoResponse(dest.IP.String(), dest.Port, reqRaw, uint32(timeout.Seconds()))
	}

	respRaw, err := ufnet.SendTCPRequest(dest.IP.String(), dest.Port, reqRaw, uint32(timeout.Seconds()))
	if err != nil {
		return err
	}

	resp := &ucloud.UMessage{}
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
	MessageBodyExtensions = proto.RegisteredExtensions((*ucloud.Body)(nil))
}
