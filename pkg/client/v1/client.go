package v1

import (
	"context"
	"encoding/json"
	"log"

	"github.com/tsetsik/ports-storage/internal/storage"
	"google.golang.org/grpc"
)

// PortsStorageClient interface for defining properties
type PortsStorageClient interface {
	UpsertPort(id int32, data map[string]interface{}) (*storage.Port, error)
}

// PortsStorage struct
type portsStorageClient struct {
	client storage.StorageServiceClient
	conn   *grpc.ClientConn
}

// NewClient used for initializing new ports client
func NewClient(target string) (PortsStorageClient, error) {
	c := new(portsStorageClient)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
		return nil, err
	}

	c.conn = conn
	c.client = storage.NewStorageServiceClient(conn)

	return c, nil
}

func (sc *portsStorageClient) UpsertPort(id int32, data map[string]interface{}) (*storage.Port, error) {
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error when marshaling data: %s", err)
		return nil, err
	}

	message := storage.Port{
		Id:   id,
		Data: string(json),
	}

	response, err := sc.client.UpsertPort(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling UpsertPort: %s", err)
		return nil, err
	}

	return response, nil
}
