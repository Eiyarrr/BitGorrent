package bittorrent

import "net/url"

// belongs to torFile (e.g. torFile.buildTrackerURL(...))
func (torFile *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(torFile.Announce)
	if err != nil {
		return "", err
	}

	return "", nil
}
