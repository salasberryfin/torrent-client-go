package client

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/salasberryfin/torrent-client-go/handshake"
	"github.com/salasberryfin/torrent-client-go/torrent"
)

// New creates a new instance of a torrent client application for the given
// .torrent file
func New(dir, file string) (Client, error) {
	filePath := path.Join(dir, file)
	t := &torrent.Torrent{
		Path: filePath,
	}
	data, err := t.Parse()
	if err != nil {
		return Client{}, err
	}
	t.Data = data
	bencodedInfo, err := t.BencodeInfo()
	if err != nil {
		return Client{}, err
	}
	t.Data.BencodedInfo = bencodedInfo
	t.Tracker, err = t.NewHTTPTracker()
	if err != nil {
		return Client{}, err
	}
	t.Hash = t.Tracker.InfoHash

	cli := Client{
		Torrent:     *t,
		Tracker:     t.Tracker,
		AnnounceURL: t.Data.Announce,
	}
	// get peer info from tracker response
	cli.trackerRequest()

	return cli, nil
}

// trackerRequest retrieves the peer info & saves it to the Client instance
func (cli *Client) trackerRequest() {
	resp := cli.Tracker.Request(cli.AnnounceURL)
	peers := torrent.GetPeers(resp.Peers)
	// keep it simple: use one peer, for now
	cli.Peer = peers[0]

	return
}

// Connect establishes a TCP connection with a peer through handshake and
// receives the bitfield message
func (cli *Client) connect() error {
	fmt.Printf("Opening connection to %v:%v\n", cli.Peer.IP, cli.Peer.Port)
	conn, err := net.DialTimeout(
		"tcp",
		fmt.Sprintf("%v:%v", cli.Peer.IP, cli.Peer.Port),
		3*time.Second,
	)
	if err != nil {
		log.Fatal("Error establishing TCP connection: ", err)
	}

	// sending handshake and validating response from peer
	h, err := handshake.NewConnection(
		conn,
		cli.Torrent.Hash,
		cli.Tracker.PeerID,
		cli.Peer,
	)
	if err != nil {
		conn.Close()
		return err
	}
	if !bytes.Equal(cli.Torrent.Hash, h.InfoHash) {
		return fmt.Errorf("Handshake response includes an invalid info hash: expected %v but got %v\n", cli.Torrent.Hash, h.InfoHash)
	}
	fmt.Printf("Valid handshake response received: %v\n", *h)

	// parsing Bitfield message after successful handshake
	bitfieldBytes, err := handshake.GetBitfield(conn)
	if err != nil {
		log.Fatalf("Failed to received the bitfield message: %v\n", err)
	}
	fmt.Printf("Bitfield received is: %v\n", bitfieldBytes)

	return nil
}

// Download starts the process of downloading to a local file from a given
// torrent filename
func Download(filename string) error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory: ", err)
	}

	client, err := New(wd, filename)
	if err != nil {
		return fmt.Errorf("Client for the given Torrent file not created: %v", err)
	}
	fmt.Printf("The peer IP is %v\n", client.Peer.IP)
	fmt.Printf("The peer port is %v\n", client.Peer.Port)

	// establish connection by first sending a TCP handhsake
	client.connect()

	return nil
}
