package network

import (
	"fmt"
	"net"

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
	} else {
		err = monitor.Network.Init(networkName)
		if err != nil {
			panic(err)
		}
		err = s.Add(&monitor.Network)
		if err != nil {
			panic(err)
		}
	}

	devices, err := monitor.scanner.ScanNetwork(monitor.Network.IP)

	if err != nil {
		panic(err)
	}

	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Printf("device found: %+v\n", device)
		deviceIP := net.ParseIP(device.IpAddress)
		err := monitor.scanner.ScanDevice(deviceIP)
		fmt.Println(err)
	}
	return monitor
}

func (m *Monitor) GetBandwidth() {
}
