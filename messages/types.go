package messages

// Message IDs for each message type
const (
	Choke         = 0
	Unchoke       = 1
	Interested    = 2
	NotInterested = 3
	Bitfield      = 5
)

// Message contains the body of the block received from a peer
type Message struct {
	Length  uint32
	ID      int
	Payload []byte
}
