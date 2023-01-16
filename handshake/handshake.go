package handshake

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

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
func NewConnection(conn net.Conn, infoHash, peerID []byte, peer torrent.Peer) (*Handshake, error) {
	handshake := Handshake{
		InfoHash: infoHash,
		PeerID:   peerID,
		Peer:     peer,
		Pstr:     pstr,
		Reserved: reserved,
	}
	msg := handshake.formatMessage()

	// tcp connection
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.SetDeadline(time.Time{})
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Failed to write to TCP connection: ", err)
	}
	// read response
	resp := read(conn)

	return resp, nil
}

func read(r io.Reader) *Handshake {
	recvBuf := make([]byte, 1)
	_, err := io.ReadFull(r, recvBuf)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Println("read timeout: ", err)
		} else {
			log.Println("read err: ", err)
		}
	}

	pstrlen := int(recvBuf[0])

	if pstrlen == 0 {
		fmt.Printf("Invalid response: pstrlen is 0\n")
	}

	rcvMsg := make([]byte, pstrlen+48)
	_, err = io.ReadFull(r, rcvMsg)
	if err != nil {
		log.Fatal("Failed to parse handshake response: ", err)
	}

	var infoHash, peerID [20]byte

	copy(infoHash[:], rcvMsg[pstrlen+8:pstrlen+8+20])
	copy(peerID[:], rcvMsg[pstrlen+8+20:])

	h := Handshake{
		Pstr:     string(rcvMsg[:pstrlen]),
		Reserved: rcvMsg[pstrlen+1 : pstrlen+8],
		InfoHash: infoHash[:],
		PeerID:   peerID[:],
	}

	return &h
}
