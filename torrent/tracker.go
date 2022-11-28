package torrent

import (
	"crypto/sha1"
	"encoding/binary"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"

	bencode "github.com/jackpal/bencode-go"
)

const peerIDBytes = 20

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
	b := make([]rune, peerIDBytes)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	randString := string(b)
	hasher := sha1.New()
	hasher.Write([]byte(randString))
	peerIDHash := hasher.Sum(nil)

	return peerIDHash
}

// NewHTTPTracker creates a new tracker for the Torrent specification
func (t *Torrent) NewHTTPTracker() (track HTTPTracker, err error) {
	track = HTTPTracker{
		InfoHash: t.computeHash(),
		PeerID:   generateRandomPeerID(),
		//Port:       net.Port,
		Port:       6881,
		Uploaded:   0,
		Downloaded: 0,
		Left:       0,
		Compact:    1,
	}

	return
}

// Parse HTTP tracker response
func (t *HTTPTracker) Parse(r *http.Response) (d TrackerResponse) {
	if r.StatusCode != 200 {
		log.Fatal("Something went wrong when sending tracker request, error:", r.StatusCode)
	}

	err := bencode.Unmarshal(r.Body, &d)
	if err != nil {
		log.Fatal("Something went wrong when parsing tracker response:", err)
	}

	return
}

// Request generates a Tracker HTTP request for the given Torrent
func (t *HTTPTracker) Request(announce string) (d TrackerResponse) {
	params := url.Values{}
	params.Set("info_hash", string(t.InfoHash))
	params.Set("peer_id", string(t.PeerID))
	params.Set("port", strconv.Itoa(t.Port))
	params.Set("uploaded", strconv.Itoa(t.Uploaded))
	params.Set("downloaded", strconv.Itoa(t.Downloaded))
	params.Set("left", strconv.Itoa(t.Downloaded))
	params.Set("compact", strconv.Itoa(1))
	params.Set("event", "started")

	client := http.Client{}
	req, err := http.NewRequest("GET", announce, nil)
	if err != nil {
		log.Fatal("Failed when building HTTP request: ", err)
	}
	req.URL.RawQuery = params.Encode()
	log.Println("Request URL: ", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed when sending tracker HTTP request: ", err)
	}
	defer resp.Body.Close()

	d = t.Parse(resp)

	if d.FailureReason != "" {
		log.Fatal("Tracker response reports error:", d.FailureReason)
	}

	return
}

// GetPeers extracts the peer info from tracker response
func GetPeers(peers string) (ips []Peer) {
	bytePeers := []byte(peers)
	numPeers := len(peers) / 6
	for x := 0; x < numPeers; x++ {
		ipBytes := bytePeers[x*6 : (x*6)+6]
		ip := net.IPv4(ipBytes[0], ipBytes[1], ipBytes[2], ipBytes[3])
		port := binary.BigEndian.Uint16(ipBytes[4:6])

		ips = append(ips, Peer{
			IP:   ip,
			Port: int(port)},
		)
	}

	return
}
