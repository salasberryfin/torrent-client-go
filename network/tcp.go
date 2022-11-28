package network

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Connect establishes a TCP connection with a peer
func (p Peer) Connect(msg []byte) (h Handshake) {
	fmt.Printf("Opening connection to %v:%v\n", p.IP, p.Port)
	conn, err := net.DialTimeout(
		"tcp",
		fmt.Sprintf("%v:%v", p.IP, p.Port),
		3*time.Second,
	)
	if err != nil {
		log.Fatal("Error establishing TCP connection: ", err)
	}
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.SetDeadline(time.Time{})

	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("Failed to write to TCP connection: ", err)
	}

	recvBuf := make([]byte, 1)
	_, err = io.ReadFull(conn, recvBuf)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Println("read timeout: ", err)
		} else {
			log.Println("read err: ", err)
		}
	}

	pstrlen := int(recvBuf[0])

	if pstrlen == 0 {
		fmt.Printf("Invalid response from %v: pstrlen is 0\n", p.IP)
	}

	rcvMsg := make([]byte, pstrlen+49)
	_, err = io.ReadFull(conn, rcvMsg)
	if err != nil {
		log.Fatal("Failed to parse handshake response: ", err)
	}

	var infoHash, peerID [20]byte

	copy(infoHash[:], rcvMsg[pstrlen+8:pstrlen+8+20])
	copy(peerID[:], rcvMsg[pstrlen+8+20:])

	h = Handshake{
		Pstr:     string(rcvMsg[:pstrlen]),
		Reserved: rcvMsg[pstrlen+1 : pstrlen+8],
		InfoHash: infoHash[:],
		PeerID:   peerID[:],
	}

	fmt.Println("Handshake response:")
	fmt.Println("\tInfoHash: ", h.InfoHash)
	fmt.Println("\tPeerID: ", h.PeerID)

	return
}
