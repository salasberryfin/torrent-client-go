package client

import (
	"bytes"
	"errors"
	"log"
	"os"
	"path"

	bencode "github.com/jackpal/bencode-go"
)

// ParseTorrentFile reads the bencoded .torrent file and parses its content
func ParseTorrentFile(dirPath string, fileName string) (Torrent, error) {
	torrentFilePath := path.Join(dirPath, fileName)
	log.Print("Opening torrent file: ", torrentFilePath)
	data, err := os.Open(torrentFilePath)
	if err != nil {
		return Torrent{}, errors.New("Reading .torrent file failed: " + err.Error())
	}

	// consider single-file only for now
	d := MetaInfo{}
	err = bencode.Unmarshal(data, &d)
	if err != nil {
		return Torrent{}, errors.New("Uncoding bencode failed: " + err.Error())
	}

	bencodedInfo := bytes.Buffer{}
	errorMarshal := bencode.Marshal(&bencodedInfo, d.Info)
	if errorMarshal != nil {
		return Torrent{}, errors.New("Bencoding Info failed: " + errorMarshal.Error())
	}
	d.BencodedInfo = bencodedInfo

	t := Torrent{Path: torrentFilePath, Data: d}
	httpTracker, err := NewTracker(t)
	if err != nil {
		return Torrent{}, errors.New("Failed to create Tracker instance: " + err.Error())
	}
	t.Tracker = *httpTracker

	return t, nil
}
