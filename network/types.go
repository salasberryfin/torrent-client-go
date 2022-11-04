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

// Payload is the interface that implements the formatMessage function for different message types
type Payload interface {
	formatMessage()
}

// RequestPayload is the payload of the message used to request a block
type RequestPayload struct {
	Index  int
	Begin  int
	Length int
}

// PiecePayload is the payload of the message that contains blocks of data
type PiecePayload struct {
	Index int
	Begin int
	Block string
}

// Handshake contains the information used for the initial TCP connection
type Handshake struct {
	Pstr     string
	Reserved []byte
	InfoHash []byte
	PeerID   []byte
}
