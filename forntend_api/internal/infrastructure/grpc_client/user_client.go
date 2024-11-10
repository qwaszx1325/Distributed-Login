package grpcclient

import (
	"context"

	"example.com/simple-login/forntend_api/internal/config"
	"example.com/simple-login/pkg/dlerr"
	"example.com/simple-login/pkg/pb/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	conn          *grpc.ClientConn
	userGrpClient user.UserServiceClient
}

func NewUserClient(cfg *config.Config) (*UserClient, error) {
	grpcAddr := cfg.UserUrl

	conn, err := grpc.NewClient(
		grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return &UserClient{}, err
	}
	userGrpc := user.NewUserServiceClient(conn)

	return &UserClient{
		conn:          conn,
		userGrpClient: userGrpc,
	}, nil
}

func (c *UserClient) close() error {
	return c.conn.Close()
}

func (a *UserClient) CreateUserProfile(ctx context.Context, req *user.CreateUserProfileRequest) (*user.CreateUserProfileResponse, *dlerr.DlError) {
	res, grpcErr := a.userGrpClient.CreateUserProfile(ctx, req)
	if grpcErr != nil {
		if err, ok := dlerr.FromGrpcErr(grpcErr); ok {
			return &user.CreateUserProfileResponse{}, err
		}
		err := dlerr.New(dlerr.InternalServerError, "can't found the dlErr from grpcErr", grpcErr)
		return &user.CreateUserProfileResponse{}, err
	}
	return res, nil
}
