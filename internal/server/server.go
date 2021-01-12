package server

import (
	"context"
	"math/rand"

	"github.com/tsetsik/ports-storage/internal/storage"
)

type server struct {
	data map[int32]string
}

// Server interface definition
type Server interface {
	UpsertPort(ctx context.Context, message *storage.Port) (*storage.Port, error)
}

// NewServer initializing new server
func NewServer() Server {
	return &server{data: map[int32]string{}}
}

// UpsertPort function for processing port record
func (s *server) UpsertPort(ctx context.Context, message *storage.Port) (*storage.Port, error) {
	id := message.Id
	if message.Id <= 0 {
		id = rand.Int31()
	}

	s.data[id] = message.Data

	return message, nil
}
