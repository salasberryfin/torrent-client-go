package client

import (
	"path"
)

// New creates a new instance of a Torrent client for the given
// .torrent file
func New(dir, file string) (t *Torrent, err error) {
	filePath := path.Join(dir, file)
	t = &Torrent{
		Path: filePath,
	}
	data, err := t.Parse()
	if err != nil {
		return
	}
	t.Data = data
	bencodedInfo, err := t.BencodeInfo()
	if err != nil {
		return
	}
	t.Data.BencodedInfo = bencodedInfo

	return
}
