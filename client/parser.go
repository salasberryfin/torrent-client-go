package client

import (
	"bytes"
	"log"
	"os"

	bencode "github.com/jackpal/bencode-go"
)

// Parse .torrent file
func (t *Torrent) Parse() (d MetaInfo, err error) {
	log.Print("Opening torrent file: ", t.Path)
	data, err := os.Open(t.Path)
	if err != nil {
		return
	}
	// consider single-file only for now
	err = bencode.Unmarshal(data, &d)
	if err != nil {
		return
	}

	return
}

// BencodeInfo parses MetaInfo.Info for later hashing
func (t *Torrent) BencodeInfo() (b bytes.Buffer, err error) {
	err = bencode.Marshal(&b, t.Data.Info)
	if err != nil {
		return
	}

	return
}
