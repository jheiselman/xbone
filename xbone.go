package xbone

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Xbox - An Xbox object
type Xbox struct {
	IP     string
	LiveID string
	Debug  bool
}

func getHexLength(payload []byte) byte {
	length := len(payload)
	b := byte(length)
	return b
}

// TurnOn - Sends a request to turn on the Xbox
func (xbox *Xbox) TurnOn() error {
	dest, err := net.ResolveUDPAddr("udp", xbox.IP+":5050")
	if err != nil {
		return fmt.Errorf("Invalid hostname")
	}

	conn, err := net.DialUDP("udp", nil, dest)
	if err != nil {
		return fmt.Errorf("Cannot connect to "+xbox.IP+": %v", err)
	}

	bLiveID := []byte(xbox.LiveID)
	payload := []byte{'\x00', getHexLength(bLiveID)}
	payload = append(payload, append(bLiveID, '\x00')...)

	header := append([]byte{'\xdd', '\x02', '\x00'}, getHexLength(payload))
	header = append(header, []byte{'\x00', '\x00'}...)

	packet := append(header, payload...)
	// Example of proper packet
	// b'\xdd\x02\x00\x13\x00\x00\x00\x10FD00A44B102D3341\x00'

	if xbox.Debug {
		log.Printf("DEBUG: Sending payload\n\t%v", packet)
	}

	for i := 0; i < 5; i++ {
		_, err := conn.Write(packet)
		if err != nil {
			return fmt.Errorf("Failed to send data: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	// This doesn't appear to actually do much
	xboxPingPayload := []byte{'\xdd', '\x00', '\x00', '\x0a', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x04', '\x00', '\x00', '\x00', '\x02'}
	_, err = conn.Write(xboxPingPayload)
	if err != nil {
		return fmt.Errorf("Xbox didn't respond to Ping: %v", err)
	}

	conn.Close()

	return nil
}
