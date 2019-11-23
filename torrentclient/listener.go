package torrentclient

import (
    "log"
    "sync"
    "net"
    "errors"
    "strconv"
)

func (torrentNetwork* Network) listenOnPort (wg *sync.WaitGroup) {
    defer wg.Done()
    log.Print("Reached listenOnPort")
}

func (torrentNetwork* Network) createListener () (int, error) {
    var err error
    var ln net.Listener

    for port := 6881; port <= 6889; port++ { // ports 6881-6889
        log.Printf("Trying to create listener on port %d", port)
        ln, err = net.Listen("tcp", ":" + strconv.Itoa(port))
        if err == nil {
            torrentNetwork.Port = port
            torrentNetwork.Listener = ln
            break
        }
    }

    if torrentNetwork.Listener == nil {
        log.Printf("Listening on port %d failed: %s", torrentNetwork.Port, err.Error())
        return torrentNetwork.Port, errors.New("Create listener operation failed: " + err.Error())
    }

    // run as a goroutine
    var wg sync.WaitGroup
    wg.Add(1)
    go torrentNetwork.listenOnPort(&wg)
    wg.Wait()

    return torrentNetwork.Port, nil
}
