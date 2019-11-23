package torrentclient

import (
    "log"
    "crypto/sha1"
    "math/rand"
	"net"
	"net/url"
	"net/http"
	"strconv"
	"sync"
)

const PEER_ID_BYTES = 20

type HttpTracker struct {
    InfoHash     []byte
    PeerId       []byte
    Port         int
    Uploaded     int
    Downloaded   int
    Left         int
    Compact      int
}

type TrackerResponse struct {
    Body         *http.Response
}

type Network struct {
    Listener    net.Listener
    Port        int
}

func computeHashes(torrent Torrent) []byte {
    // Hash bencoded Info
    info := torrent.Data.BencodedInfo
    log.Print("Computing SHA1 hash for ", torrent.Data.Info.Name)
    hasher := sha1.New()
    hasher.Write(info.Bytes())
    infoSha1 := hasher.Sum(nil)

    return infoSha1
}

func generateRandomPeerId() []byte {
	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, PEER_ID_BYTES)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
    randString := string(b)
    hasher := sha1.New()
    hasher.Write([]byte(randString))
    peerIdHash := hasher.Sum(nil)

    return peerIdHash

}

func generateTracker (torrent Torrent, wg *sync.WaitGroup) (HttpTracker, error) {
    infoHash := computeHashes(torrent)
    log.Print("Generating HTTP tracker...")
    peerId := generateRandomPeerId()
    var torrentNetwork Network
    listenPort, errListen := torrentNetwork.createListener(wg)
    if errListen != nil {
        return HttpTracker {}, errListen
    }

    return HttpTracker {InfoHash: infoHash, PeerId: peerId, Port: listenPort, Uploaded: 0, Downloaded: 0, Left: 0, Compact: 1}, nil
}

func TrackerRequest (torrent Torrent, wg *sync.WaitGroup) (TrackerResponse, error) {
    tracker, errTracker := generateTracker(torrent, wg)
    if errTracker != nil {
        return TrackerResponse {}, errTracker
    }
	queryParams := url.Values {}
	queryParams.Set("info_hash", string(tracker.InfoHash))
	queryParams.Set("peer_id", string(tracker.PeerId))
	queryParams.Set("port", strconv.Itoa(tracker.Port))
	queryParams.Set("uploaded", strconv.FormatInt(0, 10))
	queryParams.Set("downloaded", strconv.FormatInt(0, 10))
	queryParams.Set("left", strconv.FormatInt(int64(torrent.Data.Info.PieceLength), 10))
	queryParams.Set("compact", strconv.Itoa(1))
    queryParams.Set("event", "started")

    // Send HTTP Tracker request
    announce, errUrl := url.Parse(torrent.Data.Announce)
    if errUrl != nil {
        return TrackerResponse {}, errUrl
    }
    announce.RawQuery = queryParams.Encode()
    trackerResponse, errReq := http.Get(announce.String())
    if errReq != nil {
        return TrackerResponse {}, errReq
    }
    log.Print(trackerResponse)

    return TrackerResponse {Body: trackerResponse}, nil

}
