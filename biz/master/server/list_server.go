package server

import (
	"context"

	"github.com/VaalaCat/frp-panel/common"
	"github.com/VaalaCat/frp-panel/dao"
	"github.com/VaalaCat/frp-panel/models"
	"github.com/VaalaCat/frp-panel/pb"
	"github.com/samber/lo"
)

func ListServersHandler(c context.Context, req *pb.ListServersRequest) (*pb.ListServersResponse, error) {
	var (
		userInfo = common.GetUserInfo(c)
		page     = int(req.GetPage())
		pageSize = int(req.GetPageSize())
	)

	if !userInfo.Valid() {
		return &pb.ListServersResponse{
			Status: &pb.Status{Code: pb.RespCode_RESP_CODE_INVALID, Message: "invalid user"},
		}, nil
	}

	servers, err := dao.ListServers(userInfo, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &pb.ListServersResponse{
		Status: &pb.Status{Code: pb.RespCode_RESP_CODE_SUCCESS, Message: "ok"},
		Servers: lo.Map(servers, func(c *models.ServerEntity, _ int) *pb.Server {
			return &pb.Server{
				Id:     lo.ToPtr(c.ServerID),
				Config: lo.ToPtr(string(c.ConfigContent)),
			}
		}),
	}, nil
}
