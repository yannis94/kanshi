package network

import (
	"fmt"

	"github.com/yannis94/kanshi/internal/models"
	"github.com/yannis94/kanshi/internal/store"
)

type Monitor struct {
	storage store.NetworkStorage
	scanner *Scanner
	Network models.Network
}

func NewMonitor(s store.NetworkStorage, networkName string) *Monitor {
	monitor := &Monitor{
		storage: s,
		scanner: &Scanner{},
	}

	network, err := monitor.storage.Get(networkName)
	if err != nil {
		panic(err)
	}

	if network != nil {
		monitor.Network = *network
		return monitor
	}

	err = monitor.Network.Init(networkName)
	if err != nil {
		panic(err)
	}

	err = s.Add(&monitor.Network)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", monitor.Network)
	devices, err := monitor.scanner.ScanNetwork(monitor.Network.IP)

	if err != nil {
		panic(err)
	}

	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Printf("device found: %+v\n", device)
	}
	return monitor
}

func (m *Monitor) GetBandwidth() {
}
