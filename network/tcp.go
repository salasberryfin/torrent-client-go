package network

// Connect establishes a TCP connection with a peer
/*
func Connect(conn net.Conn, msg []byte, p torrent.Peer) (io.Reader, error) {
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.SetDeadline(time.Time{})

	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Failed to write to TCP connection: ", err)
	}

	return conn, nil
}
*/

/*
func Receive(conn net.Conn) (io.Reader, error) {
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.SetDeadline(time.Time{})
	resp, err := messages.Read(conn)
	if err != nil {
		return &Message{}, err
	}

	return &resp, nil
}
*/
