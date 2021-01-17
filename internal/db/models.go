package db

import "fmt"

// Model interface
type Model interface {
	Upsert() error
}

// Port db model
type Port struct {
	ID          int32
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float32
	Province    string
	Timezone    string
	Code        string
	Unlocs      []string
}

// Upsert metod in ports collection
func (p *Port) Upsert() error {
	//TODO: implementation goes here
	fmt.Println("\n\n Aide da vidim tuk neshto nai-nakraia")
	return nil
}
