package client

import (
	"crypto/sha1"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	bencode "github.com/jackpal/bencode-go"
)

const PEER_ID_BYTES = 20

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
func NewTracker(torrent Torrent) (t *HTTPTracker, err error) {
	//net, err := network.New(wg)
	t = &HTTPTracker{
		InfoHash: torrent.computeHash(),
		PeerId:   generateRandomPeerID(),
		//Port:       net.Port,
		Port:       6881,
		Uploaded:   0,
		Downloaded: 0,
		Left:       0,
		Compact:    1,
	}

	return
}

func (t *HTTPTracker) ParseResponse(r *http.Response) {
	d := TrackerResponse{}

	if r.StatusCode != 200 {
		log.Fatal("Something went wrong when sending tracker request, error:", r.StatusCode)
	}

	err := bencode.Unmarshal(r.Body, &d)
	if err != nil {
		log.Fatal("Something went wrong when parsing tracker response:", err)
	}

	fmt.Println("Tracker response: ", d)
}

// TrackerRequest generates a Tracker HTTP request for the given Torrent
func (t *Torrent) TrackerRequest() {
	tracker := t.Tracker
	params := url.Values{}
	params.Set("info_hash", string(tracker.InfoHash))
	params.Set("peer_id", string(tracker.PeerId))
	params.Set("port", strconv.Itoa(tracker.Port))
	params.Set("uploaded", strconv.Itoa(tracker.Uploaded))
	params.Set("downloaded", strconv.Itoa(tracker.Downloaded))
	params.Set("left", strconv.Itoa(tracker.Downloaded))
	params.Set("compact", strconv.Itoa(1))
	params.Set("event", "started")

	client := http.Client{}
	req, err := http.NewRequest("GET", t.Data.Announce, nil)
	if err != nil {
		log.Fatal("Failed when building HTTP request: ", err)
	}
	req.URL.RawQuery = params.Encode()
	log.Println("Request URL: ", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed when sending tracker HTTP request: ", err)
	}

	tracker.ParseResponse(resp)
}

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
