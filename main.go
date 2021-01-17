package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/tsetsik/ports-storage/internal/server"
	"github.com/tsetsik/ports-storage/internal/storage"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := server.NewServer(mongoURI)

	grpcServer := grpc.NewServer()

	storage.RegisterStorageServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(
		gracefulStop,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGKILL,
	)

	go func() {
		sig := <-gracefulStop

		// Graceful shutdown
		server.Stop(sig)

		// Shutdown grpc
		grpcServer.Stop()

		// handle it
		os.Exit(0)
	}()
}
