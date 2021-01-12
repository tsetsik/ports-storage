package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/tsetsik/ports-storage/internal/storage"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := storage.Server{}

	grpcServer := grpc.NewServer()

	storage.RegisterStorageServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
