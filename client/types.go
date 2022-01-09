package client

import "bytes"

// Torrent represent all information required by the protocol
type Torrent struct {
	Path string
	Data MetaInfo
	Hash []byte
}

// IsMultiFile checks if the decoded data corresponds to a multi-file torrent
func (i InfoDictEnvelope) IsMultiFile() (multi bool) {
	return i.Files.Length != 0
}

// MetaInfo is the struct where parsed bencoded .torrent data will be stored
type MetaInfo struct {
	Info         InfoDictEnvelope `bencode:"info"`
	Announce     string           `bencode:"announce"`
	AnnounceList [][]string       `bencode:"announce-list"`
	CreationDate int              `bencode:"creation date"`
	Comment      string           `bencode:"comment"`
	CreatedBy    string           `bencode:"created by"`
	Encoding     string           `bencode:"encoding"`
	BencodedInfo bytes.Buffer
}

// InfoDictEnvelope is the basic Info struct for both single and multi-file
type InfoDictEnvelope struct {
	Name        string    `bencode:"name"`
	Length      int       `bencode:"length,omitempty"`
	Md5Sum      string    `bencode:"md5sum,omitempty"`
	Files       FilesDict `bencode:"files,omitempty"`
	PieceLength int       `bencode:"piece length"`
	Pieces      string    `bencode:"pieces"`
	Private     int       `bencode:"private"`
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

// type PeersDict struct {
// 	PeerID string `bencode:"peer id"`
// 	IP     string `bencode:"ip"`
// 	Port   int    `bencode:"port"`
// }
