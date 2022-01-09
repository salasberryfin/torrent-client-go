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

	d := MetaInfo{}
	err = bencode.Unmarshal(data, &d)
	if err != nil {
		return Torrent{}, errors.New("Uncoding bencode failed: " + err.Error())
	}

	multiFileTorrent := d.Info.IsMultiFile()
	if multiFileTorrent {
		log.Println("Multi file torrent detected")
	}

	bencodedInfo := bytes.Buffer{}
	errorMarshal := bencode.Marshal(&bencodedInfo, d.Info)
	if errorMarshal != nil {
		return Torrent{}, errors.New("Bencoding Info failed: " + errorMarshal.Error())
	}
	d.BencodedInfo = bencodedInfo

	return Torrent{Path: torrentFilePath, Data: d}, nil
}
