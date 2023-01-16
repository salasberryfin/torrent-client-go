package handshake

import (
	"fmt"
	"net"
	"time"

	"github.com/salasberryfin/torrent-client-go/messages"
)

func GetBitfield(conn net.Conn) (Bitfield, error) {
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.SetDeadline(time.Time{})

	bitMsg, err := messages.Read(conn)
	if err != nil {
		return Bitfield{}, err
	}
	// message ID for Bitfield messages is 5
	if bitMsg.ID != messages.Bitfield {
		return Bitfield{}, fmt.Errorf("Not a bitfield message: Message ID = %d\n", bitMsg.ID)
	}

	return bitMsg.Payload, nil
}
