package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mmuoDev/wallet/gen/wallet"
	"github.com/mmuoDev/wallet/internal/app"
	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = "4444"
)

func getListener() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}
	return lis
}

func main() {
	a := app.New()
	grpcServer := grpc.NewServer()
	wallet.RegisterWalletServer(grpcServer, a)
	log.Println(fmt.Sprintf("Starting server on address: %s:%s", host, port))
	if err := grpcServer.Serve(getListener()); err != nil {
		log.Fatal("failed to serve: " + err.Error())
	}
}
