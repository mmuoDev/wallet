package main

import (
	"context"
	"log"
	"net"

	pb "github.com/mmuoDev/wallet/gen/wallet"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":50051"
)

type WalletServer struct {
	pb.UnimplementedWalletServer
}

func (w *WalletServer) CreateWallet(ctx context.Context, in *pb.CreateWalletRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (w *WalletServer) UpdateWallet(ctx context.Context, in *pb.UpdateWalletRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (w *WalletServer) RetrieveWallet(ctx context.Context, in *pb.RetrieveWalletRequest) (*pb.RetrieveWalletResponse, error) {
	return &pb.RetrieveWalletResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWalletServer(s, &WalletServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
