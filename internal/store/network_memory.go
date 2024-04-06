package store

import (
	"errors"

	"github.com/yannis94/kanshi/internal/models"
)

type NetworkMemoryStore struct {
	networks []*models.Network
}

func NewNetworkMemoryStore() *NetworkMemoryStore {
	return &NetworkMemoryStore{
		networks: make([]*models.Network, 0),
	}
}

func (s *NetworkMemoryStore) Get(ip string) (*models.Network, error) {
	for _, network := range s.networks {
		if ip == network.IP.String() {
			return network, nil
		}
	}

	return nil, nil
}

func (s *NetworkMemoryStore) Add(network *models.Network) error {
	s.networks = append(s.networks, network)
	return nil
}

func (s *NetworkMemoryStore) Update(network *models.Network) error {
	for i, n := range s.networks {
		if n.IP.Equal(network.IP) {
			s.networks[i] = network
			return nil
		}
	}

	return errors.New("network not found")
}

func (s *NetworkMemoryStore) Delete(ip string) error {
	for i, network := range s.networks {
		if network.IP.String() == ip {
			prov := append(s.networks[0:i], s.networks[i+1:]...)
			s.networks = prov
			return nil
		}
	}
	return errors.New("network not found")
}
