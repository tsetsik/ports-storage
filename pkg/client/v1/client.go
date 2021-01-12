package v1

// PortsStorageClient interface for defining properties
type PortsStorageClient interface {
	Insert() error
	Read() map[string]string
}

// PortsStorage struct
type PortsStorage struct {
}

// NewClient used for initializing new ports client
func NewClient() (*PortsStorage, error) {
	return &PortsStorage{}, nil
}
