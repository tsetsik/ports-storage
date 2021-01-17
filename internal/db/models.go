package db

// Model interface
type Model interface {
	id() int32
	collection() string
	obj() Model
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

func (p *Port) id() int32 {
	return p.ID
}

func (p *Port) collection() string {
	return "ports"
}

func (p *Port) obj() Model {
	return p
}
