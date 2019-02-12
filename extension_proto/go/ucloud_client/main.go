package main

import (
	"flag"
	"log"
	"net"
	"time"

	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud/udemo"

	"git.ucloudadmin.com/udb/v2/common"
	ufcommon "gitlab.ucloudadmin.com/udb/uframework/common"
	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
)

var serverAddress net.TCPAddr

func main() {
	parseArgs()

	var err error

	err = SendCheckUMessage(serverAddress, udemo.MessageType_value["MY_ECHO_REQUEST"], &udemo.MyEchoRequest{Id: proto.String("0"), Name: proto.String("foo")}, ufcommon.NewUUIDV4().String(), time.Second, DefaultValidater)
	if err != nil {
		log.Println(err)
	}

	err = SendCheckUMessage(serverAddress, common.GenerateGrayMessageType(udemo.MessageType_value["MY_ECHO_REQUEST"], 255), &udemo.MyEchoRequest{Id: proto.String("0"), Name: proto.String("foo")}, ufcommon.NewUUIDV4().String(), time.Second, DefaultValidater)
	if err != nil {
		log.Println(err)
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
