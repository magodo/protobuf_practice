package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"reflect"

	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud/udemo"

	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud"
	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
)

var serverAddress net.TCPAddr

func main() {
	parseArgs()

	validater := func(resp *ucloud.UMessage) error {
		respID := resp.GetHead().GetMessageType()
		respBodyExt := MessageBodyExtensions[respID]
		respBodyRaw, err := proto.GetExtension(resp.GetBody(), respBodyExt)
		if err != nil {
			return err
		}
		vResp := reflect.ValueOf(respBodyRaw).Elem()
		vId := vResp.FieldByName("Id").Elem()
		fmt.Println((vId.Interface().(string)))
		return nil
	}

	var err error
	err = SendMessage(udemo.MessageType_value["MY_ECHO_REQUEST"], &udemo.MyEchoRequest{Id: proto.String("0"), Name: proto.String("foo")}, validater)

	if err != nil {
		log.Fatal(err)
	}
}

func parseArgs() {
	ip := flag.String("h", "127.0.0.1", "server ip")
	port := flag.Int("p", 8888, "server port")
	flag.Parse()
	serverAddress = net.TCPAddr{
		IP:   net.ParseIP(*ip),
		Port: *port,
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
