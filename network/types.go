package network

import "net"

// Network defines the network configuration
type Network struct {
	Listener net.Listener
	Port     int
}

// ConnectionDetails contains the state information for each connection with a peer
type ConnectionDetails struct {
	Choked     bool
	Interested bool
}
