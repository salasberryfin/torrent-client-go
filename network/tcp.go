package network

import (
	"fmt"
	"log"
	"net"
)

// Connect establishes a TCP connection with a peer
func (p Peer) Connect(msg []byte) {
	conn, err := net.Dial(
		"tcp",
		fmt.Sprintf("%v:%v", p.IP, p.Port),
	)
	if err != nil {
		log.Fatal("Error establishing TCP connection: ", err)
	}
	_, err = conn.Write(msg)
	res, err :=
		fmt.Println("Conn: ", conn)
}
