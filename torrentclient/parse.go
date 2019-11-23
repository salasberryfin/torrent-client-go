package torrentclient

import (
    "os"
    "log"
    "bytes"
    "path"
    "errors"
    "net"
    "net/http"
    "io/ioutil"
    "strconv"

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

type TrackerResponse struct {
    //FailureReason   string      `bencode:"failure reason"`
    Interval        int         `bencode:"interval"`
    TrackerId       int         `bencode:"tracker id"`
    Complete        int         `bencode:"complete"`
    Incomplete      int         `bencode:"incomplete"`
    Peers           string      `bencode:"peers"`
}

func ParseTrackerResponse (trackerResponse *http.Response) (TrackerResponse, error) {
    d := TrackerResponse{}
    if trackerResponse.StatusCode == http.StatusOK {
        body, err := ioutil.ReadAll(trackerResponse.Body)
        if err != nil {
            return TrackerResponse {}, errors.New(err.Error())
        }
        errorBencode := bencode.Unmarshal(bytes.NewReader(body), &d)
        if errorBencode != nil {
            return TrackerResponse {}, errors.New(errorBencode.Error())
        }
    } else {
            return TrackerResponse {}, nil
    }
    defer trackerResponse.Body.Close()

    // extract Peers from 
    numPeers := len(d.Peers) / 6
    for x := 0; x < numPeers; x++ {
        ipBytes := d.Peers[x*6 : (x*6)+6]
        ip := net.IPv4(ipBytes[0], ipBytes[1], ipBytes[2], ipBytes[3])
        port, errAtoi := strconv.Atoi(strconv.Itoa(int(ipBytes[4])) + strconv.Itoa(int(ipBytes[5])))
        if errAtoi != nil {
            return TrackerResponse{}, errAtoi
        }
        log.Printf("Peer %d addr: %s:%d", x+1, ip, port)
    }


    return d, nil
}

func ParseTorrentFile (dirPath string, fileName string) (Torrent, error) {
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

