package network

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// perform a ping to the given address
func pingICMP(addr net.IP) error {
	// NOTE: Sending packets working but unable to get response

	// create udp connection
	conn, err := icmp.ListenPacket("udp4", addr.String())
	if err != nil {
		return fmt.Errorf("Unable to perform ICMP connection: %w", err)
	}

	defer conn.Close()

	pingMsg := icmp.Message{
		Type:     ipv4.ICMPTypeEcho,
		Code:     0,
		Checksum: 0,
		Body:     &icmp.Echo{ID: os.Getpid() & 0xffff, Seq: 0, Data: []byte{}},
	}

	msgBytes, err := pingMsg.Marshal([]byte(nil))

	if err != nil {
		return fmt.Errorf("Unable to marshal message: %w", err)
	}

	fmt.Printf("%s", msgBytes)

	_, err = conn.WriteTo(msgBytes, &net.IPAddr{IP: addr})

	if err != nil {
		return fmt.Errorf("Unable to write message to %s: %w", addr.String(), err)
	}
	err = conn.SetReadDeadline(time.Now().Add(time.Second * 5))

	if err != nil {
		return fmt.Errorf("Unable to set deadline on read ICMP reply: %w", err)
	}

	reply := make([]byte, 1500)
	_, _, err = conn.ReadFrom(reply)

	if err != nil {
		return fmt.Errorf("Unable to read reply: %w", err)
	}

	return nil
}

func pingTCP(addr net.IP) error {
	var (
		timeout = time.Second * 3
		target  = fmt.Sprintf("%s:%s", addr, "80")
	)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		// NOTE: Deducing that a device is connecting
		if strings.Contains(err.Error(), "connection refused") || strings.Contains(err.Error(), "permission denied") {
			return nil
		}
		return err
	}
	conn.Close()
	return nil
}
