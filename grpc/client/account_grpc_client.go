package client

import (
	"context"
	"gateway-service/common/constants"
	"gateway-service/grpc/client/account"
	"gateway-service/grpc/client/common"
	gconfig "gateway-service/grpc/config"
	"google.golang.org/grpc/metadata"
	"time"
)

type AccountClient struct {
	accountClient account.UserServiceClient
}

func NewAccountClient(accountConn *gconfig.AccountConn) *AccountClient {
	return &AccountClient{
		accountClient: account.NewUserServiceClient(accountConn.Cc),
	}
}

func (service *AccountClient) GetUser(req *account.UserRequest, md metadata.MD) (res *common.OnlyIdResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.accountClient.GetUser(ctx, req)
	return
}
