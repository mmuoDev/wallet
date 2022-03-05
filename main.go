package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mmuoDev/wallet/gen/wallet"
	"github.com/mmuoDev/wallet/internal/server"
	pg "github.com/mmuoDev/wallet/pkg/postgres"
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
	s := server.New(dbConn)
	grpcServer := grpc.NewServer()
	wallet.RegisterWalletServer(grpcServer, s)
	log.Println(fmt.Sprintf("Starting server on address: %s:%s", host, port))
	if err := grpcServer.Serve(getListener()); err != nil {
		log.Fatal("failed to serve: " + err.Error())
	}
}
