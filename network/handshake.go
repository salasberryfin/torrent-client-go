package network

import "fmt"

var (
	pstr     = "BitTorrent protocol"
	reserved = make([]byte, 8)
)

// Connect establishes a TCP connection with a peer
//func Connect(peerIP, peerPort string, msg []byte) {
//	conn, err := net.Dial(
//		"tcp",
//		fmt.Sprintf("%i:%i", peerIP, peerPort),
//	)
//	if err != nil {
//		log.Fatal("Error establishing TCP connection: ", err)
//	}
//	fmt.Println("Conn: ", conn)
//}

// formatMessage formats the content of the handshake
func (h Handshake) formatMessage() (msg []byte) {
	// It is (49+len(pstr)) bytes long
	// <pstrlen><pstr><reserved><info_hash><peer_id>
	pstrlen := len(h.Pstr)
	msg = make([]byte, pstrlen+49)
	msg[0] = byte(pstrlen)
	index := 1
	index += copy(msg[index:], pstr)
	index += copy(msg[index:], h.Reserved[:])
	index += copy(msg[index:], h.InfoHash[:])
	copy(msg[index:], h.PeerID[:])

	return
}

// NewConnection establishes a TCP connection to a peer: sends handshake and
// validates response
func NewConnection(infoHash, peerID []byte) {
	handshake := Handshake{
		InfoHash: infoHash,
		PeerID:   peerID,
		Pstr:     pstr,
		Reserved: reserved,
	}
	msg := handshake.formatMessage()
	fmt.Println("Formatted handshake:", msg)
}
