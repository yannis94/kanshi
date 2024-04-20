package network

import (
	"fmt"
	"io"
	"regexp"
	"time"

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

	for _, device := range devices {
		monitor.Network.Devices = append(monitor.Network.Devices, *device)
	}
	return monitor
}

// get bandwidth from download a file (in byte/second)
func (m *Monitor) GetBandwidth(endpoint string) (int, error) {
	if ok, err := regexp.MatchString(`^http[s]?://`, endpoint); err != nil || !ok {
		return 0, fmt.Errorf("invalid endpoint: %s, %w", endpoint, err)
	}
	// send file
	// on completion, stop timer
	start := time.Now()
	resp, err := httpGetRequest(endpoint)
	if err != nil {
		return 0, fmt.Errorf("http get request failed: %w", err)
	}
	defer resp.Body.Close()
	took := time.Since(start)

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("unable to read data: %w", err)
	}

	bandwidth := len(bytes) / int(took.Milliseconds())
	return bandwidth, nil
}
