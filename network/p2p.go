package network

import (
	"fmt"
)

var (
	pstr     = "BitTorrent protocol"
	reserved = make([]byte, 8)
)

// InitHandshake sends the first message to the peer
func InitHandshake(infoHash, peerID []byte) {
	// It is (49+len(pstr)) bytes long
	// <pstrlen><pstr><reserved><info_hash><peer_id>
	pstrlen := len(pstr)
	msg := make([]byte, pstrlen+49)
	msg[0] = byte(pstrlen)
	index := 1
	index += copy(msg[index:], pstr)
	index += copy(msg[index:], reserved)
	index += copy(msg[index:], infoHash)
	copy(msg[index:], peerID)

	fmt.Println("Handshake:", msg)

}
