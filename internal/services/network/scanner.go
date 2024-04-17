package network

import (
	"net"
	"sync"

	"github.com/yannis94/kanshi/internal/models"
)

type Scanner struct{}

func (s *Scanner) ScanNetwork(ip net.IP) ([]*models.Device, error) {
	var (
		devices   []*models.Device
		wg        sync.WaitGroup
		currentIP = ip
		mask      = ip.DefaultMask()
	)

	for i, maskAddr := range mask {
		if maskAddr == 255 {
			continue
		}

		if i != len(mask)-1 {
			continue
		}
		for j := 1; j < 256; j++ {
			target := make(net.IP, len(currentIP))
			copy(target, currentIP)

			target[i] = byte(j)
			wg.Add(1)

			go func(ip net.IP) {
				defer wg.Done()
				if err := pingTCP(ip); err != nil {
					return
				}
				device := models.Device{
					IpAddress: ip.String(),
					Connected: true,
				}
				devices = append(devices, &device)
			}(target)
		}
	}

	wg.Wait()

	return devices, nil
}

func (s *Scanner) ScanDevice(ip net.IP) error {
	return pingICMP(ip)
}
