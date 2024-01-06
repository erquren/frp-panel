package server

import (
	"github.com/VaalaCat/frp-panel/common"
	"github.com/VaalaCat/frp-panel/pb"
	"github.com/sirupsen/logrus"
)

func HandleServerMessage(req *pb.ServerMessage) *pb.ClientMessage {
	logrus.Infof("client get a server message, origin is: [%+v]", req)
	switch req.Event {
	case pb.Event_EVENT_UPDATE_FRPS:
		return common.WrapperServerMsg(req, UpdateFrpsHander)
	case pb.Event_EVENT_REMOVE_FRPS:
		return common.WrapperServerMsg(req, RemoveFrpsHandler)
	default:
	}

	return &pb.ClientMessage{
		Event: pb.Event_EVENT_ERROR,
		Data:  []byte("unknown event"),
	}
}
