package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/giavudangle/go-grpc/pb"
	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("Rrecord already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex sync.Mutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data:  make(map[string]*pb.Laptop),
		mutex: sync.Mutex{},
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// Deep copy laptop object
	other := &pb.Laptop{}

	err := copier.Copy(other, laptop)

	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other
	return nil
}
