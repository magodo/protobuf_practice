package main

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"

	"proto_foo/proto/foo"
)

type HandlerFunc func(reqBodyItf interface{}) (respBodyItf interface{})
type HandlerInfo struct {
	F          HandlerFunc
	ResponseID int32
	Timeout    time.Duration
}

// requestID -> (handler, responseID)
var HandlerInfoMap map[int32]HandlerInfo = map[int32]HandlerInfo{}

func HandleRequest(msgChan chan []byte, msg interface{}) {

	var err error
	defer func() {
		if err != nil {
			log.Fatal(err)
			close(msgChan)
		}
	}()

	// pickup handler info for this request
	req := msg.(*foo.Message)
	reqID := req.GetHeader().GetType()
	handlerInfo, ok := HandlerInfoMap[reqID]
	if !ok {
		err = errors.Errorf("no handler info registered for %d", reqID)
		return
	}
	f, respID := handlerInfo.F, handlerInfo.ResponseID

	// extract the request body
	reqBodyItf, err := proto.GetExtension(req.GetBody(), MessageBodyExtensions[reqID])
	if err != nil {
		err = errors.Wrapf(err, "failed to get extension of request: %d", reqID)
		return
	}

	/* process request body and return response body */
	respBodyItf := f(reqBodyItf)

	// prepare response
	resp := &foo.Message{
		Header: &foo.Header{Type: proto.Int32(respID)},
		Body:   &foo.Body{},
	}
	proto.SetExtension(resp.GetBody(), MessageBodyExtensions[respID], respBodyItf)

	// send back response
	buffer, err := proto.Marshal(resp)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal response")
		return
	}
	msgChan <- buffer

	return
}
