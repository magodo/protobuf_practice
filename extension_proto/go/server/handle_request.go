package main

import (
	"log"
	"time"

	"git.ucloudadmin.com/udb/v2/common"
	"github.com/pkg/errors"
	ufmessage "gitlab.ucloudadmin.com/udb/uframework/message"
	"gitlab.ucloudadmin.com/udb/uframework/message/protobuf/proto"

	"gitlab.ucloudadmin.com/udb/proto_go/proto/ucloud"
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
	req := msg.(*ucloud.UMessage)
	reqID := req.GetHead().GetMessageType()
	handlerInfo, ok := HandlerInfoMap[reqID]
	if !ok {
		err = errors.Errorf("no handler info registered for %d", reqID)
		return
	}
	f, respID := handlerInfo.F, handlerInfo.ResponseID

	validReqID, _ := common.GetMessageGrayInfo(reqID)
	// extract the request body
	reqBodyItf, err := proto.GetExtension(req.GetBody(), MessageBodyExtensions[validReqID])
	if err != nil {
		err = errors.Wrapf(err, "failed to get extension of request: %d", validReqID)
		return
	}

	/* process request body and return response body */
	respBodyItf := f(reqBodyItf)

	// prepare response
	resp, err := ufmessage.MakeResponse(req, respID)
	if err != nil {
		err = errors.Wrap(err, "failed to make response")
		return
	}
	err = proto.SetExtension(resp.GetBody(), MessageBodyExtensions[respID], respBodyItf)
	if err != nil {
		err = errors.Wrap(err, "failed to set extension")
		return
	}

	// send back response
	buffer, err := proto.Marshal(resp)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal response")
		return
	}
	msgChan <- buffer

	return
}
