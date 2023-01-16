package network

/*
All messages other than the handshake follow the format <length prefix><message ID><payload>:
	- length prefix: 4byte big-endian value
	- message ID: single decimal byte
	- payload: message dependent

Types of messages:
- keep-alive: 		length_prefix=0000
- choke: 			length_prefix=0001, id=0
- unchoke: 			length_prefix=0001, id=1
- interested:		length_prefix=0001, id=2
- not_interested:	length_prefix=0001, id=3
- have:				length_prefix=0005, piece_index
- bitfield:			length_prefix=0001+len(bitfield), id=5, bitfield
	optional
- request:			length_prefix=0013, id=6, index, begin, length
	used to request a block, contains following:
    - index: integer specifying the zero-based piece index
    - begin: integer specifying the zero-based byte offset within the piece
    - length: integer specifying the requested length.
- piece:			length_prefix=0009+len(block), id=7, index, begin, block
	payload contains following:
    - index: integer specifying the zero-based piece index
    - begin: integer specifying the zero-based byte offset within the piece
    - length: integer specifying the requested length.
- cancel:			length_prefix=0013, id=8, index, begin, length
- port:				length_prefix=0003, id=9, liste_port
*/

func (payload RequestPayload) formatMessage() (b []byte) {
	//lenPrefix := []byte{0013}
	//messageID := 6
	//msg = make([]byte, len(lenPrefix)+1+len([]byte{payload}))

	return
}

// sendMessage sends a request for a block
//func (payload RequestPayload) sendMessage() {
//	lengthPrefix := 0013
//	id := 6
//}
