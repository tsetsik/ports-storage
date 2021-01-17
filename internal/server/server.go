package server

import (
	"context"
	"log"

	"github.com/tsetsik/ports-storage/internal/db"
	"github.com/tsetsik/ports-storage/internal/storage"
)

type server struct {
	data map[int32]string
	db   db.DB
}

// Server interface definition
type Server interface {
	UpsertPort(ctx context.Context, message *storage.Port) (*storage.Port, error)
}

// NewServer initializing new server
func NewServer() Server {
	db, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	return &server{data: map[int32]string{}, db: db}
}

// UpsertPort function for processing port record
func (s *server) UpsertPort(ctx context.Context, message *storage.Port) (*storage.Port, error) {
	err := s.db.Upsert(toPortDbModel(message))
	if err != nil {
		return nil, err
	}

	return message, nil
}

func toPortDbModel(sp *storage.Port) *db.Port {
	model := &db.Port{
		ID:          sp.Id,
		Name:        sp.Name,
		City:        sp.City,
		Country:     sp.Country,
		Alias:       sp.Alias,
		Regions:     sp.Regions,
		Coordinates: sp.Coordinates,
		Province:    sp.Province,
		Timezone:    sp.Timezone,
		Unlocs:      sp.Unlocs,
		Code:        sp.Code,
	}

	return model
}
