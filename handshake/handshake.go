package handshake

import (
	"fmt"

	"github.com/salasberryfin/torrent-client-go/network"
	"github.com/salasberryfin/torrent-client-go/torrent"
)

var (
	pstr     = "BitTorrent protocol"
	reserved = make([]byte, 8)
)

// formatMessage formats the content of the handshake
func (h Handshake) formatMessage() (msg []byte) {
	// It is (49+len(pstr)) bytes long
	// <pstrlen><pstr><reserved><info_hash><peer_id>
	fmt.Println("Handshake request:")
	fmt.Println("\tInfoHash: ", h.InfoHash)
	fmt.Println("\tPeerID: ", h.PeerID)

	pstrlen := len(h.Pstr)
	msg = make([]byte, pstrlen+49)
	msg[0] = byte(pstrlen)
	index := 1
	index += copy(msg[index:], pstr)
	index += copy(msg[index:], h.Reserved[:])
	index += copy(msg[index:], h.InfoHash[:])
	index += copy(msg[index:], h.PeerID[:])

	return
}

// NewConnection establishes a TCP connection to a peer: sends handshake and
// validates response
func NewConnection(infoHash, peerID []byte, peer torrent.Peer) {
	handshake := Handshake{
		InfoHash: infoHash,
		PeerID:   peerID,
		Peer:     peer,
		Pstr:     pstr,
		Reserved: reserved,
	}
	msg := handshake.formatMessage()
	network.Peer{
		IP:   peer.IP,
		Port: peer.Port,
	}.Connect(msg)
}
