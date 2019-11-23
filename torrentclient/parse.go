package torrentclient

import (
    "os"
    "log"
    "bytes"
    "path"
    "errors"

    bencode "github.com/jackpal/bencode-go"
)

type InfoMap struct {
    Name string `bencode:"name"`
    PieceLength int `bencode:"piece length"`
    Pieces string `bencode:"pieces"`
}

type MetaInfo struct {
    Announce     string `bencode:"announce"`
    AnnounceList [][]string `bencode:"announce-list"`
    Encoding     string `bencode:"encoding"`
    Info         InfoMap `bencode:"info"`
    BencodedInfo bytes.Buffer
}

type Torrent struct {
    Path string
    Data MetaInfo
    Hash []byte
}

func Parse(dirPath string, fileName string) (Torrent, error) {
	torrentFilePath := path.Join(dirPath, fileName)
    log.Print("Opening torrent file: ", torrentFilePath)
    data, err := os.Open(torrentFilePath)
    if err != nil {
        return Torrent{}, errors.New("Reading .torrent file failed: " + err.Error())
    }

    d := MetaInfo{}
    errorBencode := bencode.Unmarshal(data, &d)
    if errorBencode != nil {
        return Torrent{}, errors.New("Uncoding bencode failed: " + errorBencode.Error())
    }

    bencodedInfo := bytes.Buffer{}
    errorMarshal := bencode.Marshal(&bencodedInfo, d.Info)
    if errorMarshal != nil {
        return Torrent{}, errors.New("Bencoding Info failed: " + errorMarshal.Error())
    }
    d.BencodedInfo = bencodedInfo

    return Torrent{Path: torrentFilePath, Data: d}, nil
}

