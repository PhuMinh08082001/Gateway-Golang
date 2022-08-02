package client

import (
	"context"
	"gateway-service/client/account"
	"gateway-service/client/common"
	"gateway-service/client/connect"
	"gateway-service/common/constants"
	"google.golang.org/grpc/metadata"
	"time"
)

type AccountClient struct {
	accountClient account.UserServiceClient
}

func NewAccountClient(accountConn *connect.AccountConn) *AccountClient {
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
