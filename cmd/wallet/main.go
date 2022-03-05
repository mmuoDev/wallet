package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/mmuoDev/core-proto/gen/wallet"
	"github.com/mmuoDev/wallet/internal/db"
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

//migratePostgre migrates postgres migrations
func migratePostgre(dbConn *pg.Connector) {
	driver, err := postgres.WithInstance(dbConn.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to connect with error %s", err)
	}
	db.MigrateDB(dbConn.DB, driver, "postgre")
}

//getPostgreConfig returns postgres conn configs
func getPostgreConfig() pg.Config {
	cfg := pg.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return cfg
}

func main() {
	dbConn, err := pg.NewConnector(getPostgreConfig())
	if err != nil {
		log.Fatal(err)
	}
	migratePostgre(dbConn)
	s := server.New(dbConn)
	grpcServer := grpc.NewServer()
	wallet.RegisterWalletServer(grpcServer, s)
	log.Println(fmt.Sprintf("Starting server on address: %s:%s", host, port))
	if err := grpcServer.Serve(getListener()); err != nil {
		log.Fatal("failed to serve: " + err.Error())
	}
}
