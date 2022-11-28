package client

import (
	"path"

	"github.com/salasberryfin/torrent-client-go/handshake"
	"github.com/salasberryfin/torrent-client-go/torrent"
)

// New creates a new instance of a torrent client application for the given
// .torrent file
func New(dir, file string) (cli Client, err error) {
	filePath := path.Join(dir, file)
	t := &torrent.Torrent{
		Path: filePath,
	}
	data, err := t.Parse()
	if err != nil {
		return
	}
	t.Data = data
	bencodedInfo, err := t.BencodeInfo()
	if err != nil {
		return
	}
	t.Data.BencodedInfo = bencodedInfo
	t.Tracker, err = t.NewHTTPTracker()
	if err != nil {
		return
	}
	t.Hash = t.Tracker.InfoHash

	cli = Client{
		Torrent:     *t,
		Tracker:     t.Tracker,
		AnnounceURL: t.Data.Announce,
	}
	// get peer info from tracker response
	cli.trackerRequest()

	return
}

// trackerRequest retrieves the peer info & saves it to the Client instance
func (cli *Client) trackerRequest() {
	resp := cli.Tracker.Request(cli.AnnounceURL)
	peers := torrent.GetPeers(resp.Peers)
	// keep it simple: use one peer, for now
	cli.Peer = peers[0]

	return
}

// Connect establishes a TCP connection with a peer through handshake
func (cli *Client) Connect() {
	handshake.NewConnection(
		cli.Torrent.Hash,
		cli.Tracker.PeerID,
		cli.Peer,
	)
}
