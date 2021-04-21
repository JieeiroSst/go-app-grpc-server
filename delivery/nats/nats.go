package nats

import (
	"github.com/nats-io/nats.go"
)

type NatService struct {
	service *nats.Conn
}

func NewSerice(service *nats.Conn) *NatService {
	return &NatService{
		service: service,
	}
}

func (n *NatService) Reciver() error {
	
}