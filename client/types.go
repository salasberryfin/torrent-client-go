package client

import "bytes"

// Torrent represent all information required by the protocol
type Torrent struct {
	Path    string
	Data    MetaInfo
	Hash    []byte
	Tracker HTTPTracker
}

// MetaInfo is the struct where parsed bencoded .torrent data will be stored
type MetaInfo struct {
	//Info         InfoDictEnvelope `bencode:"info"`
	Info         TorrentInfo `bencode:"info"`
	Announce     string      `bencode:"announce"`
	AnnounceList [][]string  `bencode:"announce-list"`
	CreationDate int         `bencode:"creation date"`
	Comment      string      `bencode:"comment"`
	CreatedBy    string      `bencode:"created by"`
	Encoding     string      `bencode:"encoding"`
	BencodedInfo bytes.Buffer
}

// TorrentInfo is the info dictionary for single file torrents
type TorrentInfo struct {
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

// HTTPTracker
type HTTPTracker struct {
	InfoHash   []byte
	PeerId     []byte
	Port       int
	Uploaded   int
	Downloaded int
	Left       int
	Compact    int
}

// FilesDict is only present when multi-file
type FilesDict struct {
	Length int      `bencode:"length"`
	Md5Sum string   `bencode:"md5sum"`
	Path   []string `bencode:"path"`
}

// TrackerResponse is a bencoded dict
type TrackerResponse struct {
	FailureReason  string `bencode:"failure reason"`
	WarningMessage string `bencode:"warning message"`
	Interval       int    `bencode:"interval"`
	MinInterval    int    `bencode:"min interval"`
	TrackerID      string `bencode:"tracker id"`
	Complete       int    `bencode:"complete"`
	Incomplete     int    `bencode:"incomplete"`
	Peers          string `bencode:"peers"`
	//Peers PeersDict `bencode:"peers"`
}

type PeersDict struct {
	PeerID string `bencode:"peer id"`
	IP     string `bencode:"ip"`
	Port   int    `bencode:"port"`
}
