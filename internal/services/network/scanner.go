package network

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/yannis94/kanshi/internal/models"
)

type Scanner struct{}

func (s *Scanner) ScanNetwork(ip net.IP) ([]*models.Device, error) {
	var (
		devices   []*models.Device
		currentIP = ip
		mask      = ip.DefaultMask()
		ch        = make(chan *models.Device)
	)

	fmt.Printf("scanner > %v\n", currentIP)
	// range over network mask
	for i, maskAddr := range mask {
		if maskAddr == 255 {
			continue
		}

		if i == len(mask)-1 {
			for j := 1; j < 256; j++ {
				currentIP[i] = byte(j)
				currentIP := currentIP
				fmt.Printf("%v scanning...\n", currentIP)
				go scanAddress(currentIP, ch)
			}
			close(ch)
		}

	}

	for {
		select {
		case device, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return devices, nil
			}
			fmt.Println("device found: ", device)
			devices = append(devices, device)
		}
	}
}

func (s *Scanner) ScanDevice() {
}

func scanAddress(addr net.IP, ch chan *models.Device) {
	var (
		timeout = time.Second * 5
		target  = fmt.Sprintf("%s:%s", addr, "80")
	)
	fmt.Printf("scanAddress > %v\n", addr)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") || strings.Contains(err.Error(), "permission denied") {
			device := &models.Device{
				IpAddress: addr.String(),
			}
			ch <- device
			<-ch
		}
		return
	}
	fmt.Printf("%v > device found: Conn> %v\n", addr, conn)
	defer conn.Close()
	device := &models.Device{
		IpAddress: addr.String(),
	}
	ch <- device
	<-ch
}
