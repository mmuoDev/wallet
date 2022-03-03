package app

import (
	"context"
	"log"
	"os"

	pb "github.com/mmuoDev/wallet/gen/wallet"
	"github.com/mmuoDev/wallet/internal/db"
	"github.com/mmuoDev/wallet/internal/workflow"
	pg "github.com/mmuoDev/wallet/pkg/postgres"
	"google.golang.org/protobuf/types/known/emptypb"
	w "github.com/mmuoDev/wallet/gen/wallet"
)

//App represents the app
type App struct {
	Option Option
	w.UnimplementedWalletServer
}

//Option represents optional args
type Option struct {
	CreateWallet db.CreateWalletFunc
}

type OptionalArg func(oa *Option)

func New(opts ...OptionalArg) App {
	cfg := pg.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	dbConn, err := pg.NewConnector(cfg)
	if err != nil {
		log.Fatal(err)
	}
	o := Option{
		CreateWallet: db.CreateWallet(*dbConn),
	}
	for _, option := range opts {
		option(&o)
	}
	return App{Option: o}
}

//CreateWallet implements grpc its method
func (a App) CreateWallet(ctx context.Context, in *pb.CreateWalletRequest) (*emptypb.Empty, error) {
	create := workflow.CreateWallet(a.Option.CreateWallet)
	if err := create(*in); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (a App) UpdateWallet(ctx context.Context, in *pb.UpdateWalletRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (a App) RetrieveWallet(ctx context.Context, in *pb.RetrieveWalletRequest) (*pb.RetrieveWalletResponse, error) {
	return &pb.RetrieveWalletResponse{}, nil
}
