package server

import (
	"context"

	pb "github.com/mmuoDev/core-proto/gen/wallet"
	"github.com/mmuoDev/wallet/internal/db"
	"github.com/mmuoDev/wallet/internal/workflow"
	pg "github.com/mmuoDev/wallet/pkg/postgres"
	"google.golang.org/protobuf/types/known/emptypb"
)

//Server represents the wallet's server
type Server struct {
	Option Option
	pb.UnimplementedWalletServer
}

//Option represents optional args
type Option struct {
	CreateWallet   db.CreateWalletFunc
	RetrieveWallet db.RetrieveWalletByAccountIdFunc
	UpdateWallet   db.UpdateWalletFunc
}

type OptionalArg func(oa *Option)

func New(dbConnector *pg.Connector, opts ...OptionalArg) Server {

	o := Option{
		CreateWallet:   db.CreateWallet(*dbConnector),
		RetrieveWallet: db.RetrieveWalletByAccountId(*dbConnector),
		UpdateWallet:   db.UpdateWallet(*dbConnector),
	}
	for _, option := range opts {
		option(&o)
	}
	return Server{Option: o}
}

//CreateWallet implements its grpc method
func (s Server) CreateWallet(ctx context.Context, in *pb.CreateWalletRequest) (*pb.CreateWalletResponse, error) {
	create := workflow.CreateWallet(s.Option.CreateWallet)
	res, err := create(*in)
	if err != nil {
		return &pb.CreateWalletResponse{}, err
	}
	return &res, nil
}

//UpdateWallet implements its grpc method
func (s Server) UpdateWallet(ctx context.Context, in *pb.UpdateWalletRequest) (*emptypb.Empty, error) {
	update := workflow.UpdateWallet(s.Option.UpdateWallet)
	if err := update(*in); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

//RetrieveWallet implements its grpc method
func (s Server) RetrieveWallet(ctx context.Context, in *pb.RetrieveWalletRequest) (*pb.RetrieveWalletResponse, error) {
	retrieve := workflow.RetrieveWallet(s.Option.RetrieveWallet)
	res, err := retrieve(*in)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
