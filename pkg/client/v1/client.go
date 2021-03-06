package v1

import (
	"context"
	"log"

	"github.com/tsetsik/ports-storage/internal/storage"
	"google.golang.org/grpc"
)

// PortsStorageClient interface for defining properties
type PortsStorageClient interface {
	UpsertPort(p *Port) (*Port, error)
}

// PortsStorage struct
type portsStorageClient struct {
	client storage.StorageServiceClient
	conn   *grpc.ClientConn
}

// Port struct for usage with this client
type Port struct {
	ID          int32     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	City        string    `json:"city,omitempty"`
	Country     string    `json:"country,omitempty"`
	Alias       []string  `json:"alias,omitempty"`
	Regions     []string  `json:"regions,omitempty"`
	Coordinates []float32 `json:"coordinates,omitempty"`
	Province    string    `json:"province,omitempty"`
	Timezone    string    `json:"timezone,omitempty"`
	Code        string    `json:"code,omitempty"`
	Unlocs      []string  `json:"unlocks,omitempty"`
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

func (sc *portsStorageClient) UpsertPort(port *Port) (*Port, error) {
	message := toStoragePort(port)

	response, err := sc.client.UpsertPort(context.Background(), message)
	if err != nil {
		log.Printf("Error when calling UpsertPort: %s", err)
		return nil, err
	}

	return toClientPort(response), nil
}

func toStoragePort(port *Port) *storage.Port {
	return &storage.Port{
		Id:          port.ID,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}

func toClientPort(port *storage.Port) *Port {
	p := &Port{
		ID:          port.Id,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}

	return p
}
