package torrentclient

import (
    "log"
    "crypto/sha1"
    "math/rand"
	"net/url"
	"net"
)

/*
info_hash 
peer_id
port
uploaded
downloaded
left
compact: Setting this to 1 indicates that the client accepts a compact response. The peers list is replaced by a peers string with 6 bytes per peer. The first four bytes are the host (in network byte order), the last two bytes are the port (again in network byte order). It should be noted that some trackers only support compact responses (for saving bandwidth) and either refuse requests without "compact=1" or simply send a compact response unless the request contains "compact=0" (in which case they will refuse the request.)
event: If specified, must be one of started, completed, stopped, (or empty which is the same as not being specified). If not specified, then this request is one performed at regular intervals.
started: The first request to the tracker must include the event key with this value.
stopped: Must be sent to the tracker if the client is shutting down gracefully.
completed: Must be sent to the tracker when the download completes. However, must not be sent if the download was already 100% complete when the client started. Presumably, this is to allow the tracker to increment the "completed downloads" metric based solely on this event.
*/

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

func GenerateTracker (torrent Torrent) (HttpTracker, error) {
	infoHash := computeHashes(torrent)
    log.Print("Generating HTTP tracker...")
    //log.Print("SHA1 Info Hash: ", infoHash)
    peerId := generateRandomPeerId()
    //log.Print("Random PeerId SHA1 hash: ", peerId)
    var torrentNetwork Network
    port, errListen := torrentNetwork.createListener()
    if errListen != nil {
        return HttpTracker {}, errListen
    }
	log.Print("Listening on port: ", port)
	queryParams := url.Values {}
	queryParams.Set("info_hash", string(infoHash))
	queryParams.Set("peer_id", string(peerId))
	log.Print("Query parameters: ", queryParams.Encode())

	return HttpTracker {InfoHash: infoHash, PeerId: peerId}, nil
}
