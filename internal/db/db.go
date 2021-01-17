package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB interface
type DB interface {
	Upsert(m Model) error
}

type db struct {
	conn *mongo.Client
}

// NewConnection initiating new db connection
func NewConnection() (DB, error) {
	moongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://test-user:test-password@localhost:27017"))
	if err != nil {
		return nil, err
	}

	return &db{conn: moongoClient}, nil
}

// Upsert is
func (db *db) Upsert(m Model) error {
	return m.Upsert()
}
