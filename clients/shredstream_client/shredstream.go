package shredstream_client

import (
	"crypto/tls"

	"github.com/17535250630/jito_cli/pkg"
	"github.com/17535250630/jito_cli/proto"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	GrpcConn *grpc.ClientConn
	RpcConn  *rpc.Client

	ShredstreamClient proto.ShredstreamClient

	Auth *pkg.AuthenticationService
}

func NewShredstreamClient(grpcDialURL string, rpcClient *rpc.Client, privateKey solana.PrivateKey) (*Client, error) {
	conn, err := grpc.Dial(grpcDialURL, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		return nil, err
	}

	shredstreamService := proto.NewShredstreamClient(conn)
	authService := pkg.NewAuthenticationService(conn, privateKey)
	if err = authService.AuthenticateAndRefresh(proto.Role_SHREDSTREAM_SUBSCRIBER); err != nil {
		return nil, err
	}

	return &Client{
		GrpcConn:          conn,
		RpcConn:           rpcClient,
		ShredstreamClient: shredstreamService,
		Auth:              authService,
	}, nil
}

func (c *Client) SendHeartbeat(count uint64, opts ...grpc.CallOption) (*proto.HeartbeatResponse, error) {
	return c.ShredstreamClient.SendHeartbeat(c.Auth.GrpcCtx, &proto.Heartbeat{Count: count}, opts...)
}
