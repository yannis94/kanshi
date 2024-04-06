package store

import "github.com/yannis94/kanshi/internal/models"

type NetworkStorage interface {
	Get(ip string) (*models.Network, error)
	Add(n *models.Network) error
	Update(n *models.Network) error
	Delete(ip string) error
}
