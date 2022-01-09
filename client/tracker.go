package client

import (
	"crypto/sha1"
	"log"
	"math/rand"
	"sync"

	"github.com/salasberryfin/torrent-client-go/network"
)

const PEER_ID_BYTES = 20

type HTTPTracker struct {
	InfoHash   []byte
	PeerId     []byte
	Port       int
	Uploaded   int
	Downloaded int
	Left       int
	Compact    int
}

func (t *Torrent) computeHash() []byte {
	info := t.Data.BencodedInfo
	log.Print("Computing SHA1 hash for ", t.Data.Info.Name)
	hasher := sha1.New()
	hasher.Write(info.Bytes())
	infoSha1 := hasher.Sum(nil)

	return infoSha1
}

func generateRandomPeerID() []byte {
	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, PEER_ID_BYTES)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	randString := string(b)
	hasher := sha1.New()
	hasher.Write([]byte(randString))
	peerIDHash := hasher.Sum(nil)

	return peerIDHash
}

// NewTracker creates a new instance of HTTPTracker
func NewTracker(torrent Torrent, wg *sync.WaitGroup) (t *HTTPTracker, err error) {
	net, err := network.New(wg)
	t = &HTTPTracker{
		InfoHash:   torrent.computeHash(),
		PeerId:     generateRandomPeerID(),
		Port:       net.Port,
		Uploaded:   0,
		Downloaded: 0,
		Left:       0,
		Compact:    1,
	}

	return
}

// func generateTracker(torrent Torrent, wg *sync.WaitGroup) (HttpTracker, error) {
// 	// infoHash := computeHashes(torrent)
// 	log.Print("Generating HTTP tracker...")
// 	peerId := generateRandomPeerID()
// 	var torrentNetwork Network
// 	listenPort, errListen := torrentNetwork.createListener(wg)
// 	if errListen != nil {
// 		return HttpTracker{}, errListen
// 	}
//
// 	return HttpTracker{InfoHash: infoHash, PeerId: peerId, Port: listenPort, Uploaded: 0, Downloaded: 0, Left: 0, Compact: 1}, nil
// }
//
// // TrackerRequest
// func TrackerRequest(torrent Torrent, wg *sync.WaitGroup) (*http.Response, error) {
// 	tracker, errTracker := generateTracker(torrent, wg)
// 	if errTracker != nil {
// 		return nil, errTracker
// 	}
// 	queryParams := url.Values{}
// 	queryParams.Set("info_hash", string(tracker.InfoHash))
// 	queryParams.Set("peer_id", string(tracker.PeerId))
// 	queryParams.Set("port", strconv.Itoa(tracker.Port))
// 	queryParams.Set("uploaded", strconv.FormatInt(0, 10))
// 	queryParams.Set("downloaded", strconv.FormatInt(0, 10))
// 	queryParams.Set("left", strconv.FormatInt(int64(torrent.Data.Info.PieceLength), 10))
// 	queryParams.Set("compact", strconv.Itoa(1))
// 	queryParams.Set("event", "started")
//
// 	// Send HTTP Tracker request
// 	announce, errUrl := url.Parse(torrent.Data.Announce)
// 	if errUrl != nil {
// 		return nil, errUrl
// 	}
// 	announce.RawQuery = queryParams.Encode()
// 	trackerResponse, errReq := http.Get(announce.String())
// 	if errReq != nil {
// 		return nil, errReq
// 	}
//
// 	return trackerResponse, nil
// }
//
// func ParseTrackerResponse(trackerResponse *http.Response) (TrackerResponse, error) {
// 	d := TrackerResponse{}
// 	if trackerResponse.StatusCode == http.StatusOK {
// 		body, err := ioutil.ReadAll(trackerResponse.Body)
// 		if err != nil {
// 			return TrackerResponse{}, errors.New(err.Error())
// 		}
// 		errorBencode := bencode.Unmarshal(bytes.NewReader(body), &d)
// 		if errorBencode != nil {
// 			return TrackerResponse{}, errors.New(errorBencode.Error())
// 		}
// 	} else {
// 		return TrackerResponse{}, nil
// 	}
// 	defer trackerResponse.Body.Close()
//
// 	// extract Peers from
// 	numPeers := len(d.Peers) / 6
// 	for x := 0; x < numPeers; x++ {
// 		ipBytes := d.Peers[x*6 : (x*6)+6]
// 		ip := net.IPv4(ipBytes[0], ipBytes[1], ipBytes[2], ipBytes[3])
// 		port, errAtoi := strconv.Atoi(strconv.Itoa(int(ipBytes[4])) + strconv.Itoa(int(ipBytes[5])))
// 		if errAtoi != nil {
// 			return TrackerResponse{}, errAtoi
// 		}
// 		log.Printf("Peer %d addr: %s:%d", x+1, ip, port)
// 	}
//
// 	return d, nil
// }
