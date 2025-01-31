package client

import (
	"context"

	"github.com/VaalaCat/frp-panel/common"
	"github.com/VaalaCat/frp-panel/dao"
	"github.com/VaalaCat/frp-panel/models"
	"github.com/VaalaCat/frp-panel/pb"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

func ListClientsHandler(ctx context.Context, req *pb.ListClientsRequest) (*pb.ListClientsResponse, error) {
	logrus.Infof("list client, req: [%+v]", req)

	var (
		userInfo = common.GetUserInfo(ctx)
	)

	if !userInfo.Valid() {
		return &pb.ListClientsResponse{
			Status: &pb.Status{Code: pb.RespCode_RESP_CODE_INVALID, Message: "invalid user"},
		}, nil
	}

	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	clients, err := dao.ListClients(userInfo, page, pageSize)
	if err != nil {
		return nil, err
	}

	clientCounts, err := dao.CountClients(userInfo)
	if err != nil {
		return nil, err
	}

	respClients := lo.Map(clients, func(c *models.ClientEntity, _ int) *pb.Client {
		return &pb.Client{
			Id:       lo.ToPtr(c.ClientID),
			Secret:   lo.ToPtr(c.ConnectSecret),
			Config:   lo.ToPtr(string(c.ConfigContent)),
			ServerId: lo.ToPtr(c.ServerID),
		}
	})

	return &pb.ListClientsResponse{
		Status:  &pb.Status{Code: pb.RespCode_RESP_CODE_SUCCESS, Message: "ok"},
		Clients: respClients,
		Total:   lo.ToPtr(int32(clientCounts)),
	}, nil
}
