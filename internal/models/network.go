package models

import (
	"errors"
	"net"
)

type Network struct {
	Name    string
	IP      net.IP
	Devices []Device
	CIDR    int
}

func NewNetwork() *Network {
	return &Network{}
}

func (n *Network) Init(name string) error {
	localAddresses, err := net.InterfaceAddrs()
	if err != nil {
		return errors.New("unable to get local interface addresses")
	}

	for _, addr := range localAddresses {
		switch ip := addr.(type) {
		case *net.IPNet:
			if ip.IP.To4() != nil && ip.IP.IsPrivate() {
				n.IP = ip.IP.Mask(ip.IP.DefaultMask())
				mask, _ := ip.IP.DefaultMask().Size()
				n.CIDR = mask
			}
		}
	}
	n.Name = name

	return nil
}
