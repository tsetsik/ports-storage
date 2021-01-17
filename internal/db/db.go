package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB interface
type DB interface {
	Upsert(m Model) error
	Shutdown()
}

type db struct {
	conn *mongo.Client
}

const databse = "ports"

// NewConnection initiating new db connection
func NewConnection() (DB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://test-user:test-password@localhost:27017/ports"))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	return &db{conn: client}, nil
}

// Upsert a mongo db record
func (db *db) Upsert(m Model) error {
	collection := db.conn.Database(databse).Collection(m.collection())

	opts := options.Replace().SetUpsert(true)
	filter := bson.D{primitive.E{Key: "_id", Value: m.id()}}
	_, err := collection.ReplaceOne(context.Background(), filter, m.obj(), opts)

	if err != nil {
		return err
	}

	return nil
}

func (db *db) Shutdown() {
	db.conn.Disconnect(context.Background())
}
