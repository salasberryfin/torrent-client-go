package network

import "net"

// Network defines the network configuration
type Network struct {
	Listener net.Listener
	Port     int
}
