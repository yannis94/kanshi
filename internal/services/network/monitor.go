package network

import (
	"fmt"

	"github.com/yannis94/kanshi/internal/models"
	"github.com/yannis94/kanshi/internal/store"
)

type Monitor struct {
	storage store.NetworkStorage
	Network models.Network
}

func NewMonitor(s store.NetworkStorage, networkName string) *Monitor {
	monitor := &Monitor{
		storage: s,
	}

	network, err := monitor.storage.Get(networkName)
	if err != nil {
		panic(err)
	}

	if network != nil {
		monitor.Network = *network
		return monitor
	}

	err = monitor.Network.Init()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", monitor)

	return monitor
}
