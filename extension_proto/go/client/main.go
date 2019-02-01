package main

import (
	"flag"
	"log"
	"net"
	"proto_foo/proto/foo/ext"

	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"
)

var serverAddress net.TCPAddr

func main() {
	parseArgs()
	err := SendMessage("foo.ext.echo_request", &ext.EchoRequest{Msg: proto.String("Hello")}, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = SendMessage("foo.ext.foo1_request", &ext.FooRequest{}, DefaultValidater)
	if err != nil {
		log.Fatal(err)
	}
	err = SendMessage("foo.ext.foo2_request", &ext.FooRequest{}, DefaultValidater)
	if err != nil {
		log.Fatal(err)
	}
}

func parseArgs() {
	ip := flag.String("h", "127.0.0.1", "server ip")
	port := flag.Int("p", 8888, "server port")
	serverAddress = net.TCPAddr{
		IP:   net.ParseIP(*ip),
		Port: *port,
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
