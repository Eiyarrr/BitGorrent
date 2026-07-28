package bittorrent

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/url"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

// parses peer IPs & ports from a buffer
func Unmarshal(peersBin []byte) ([]Peer, error) {
	const peerSize = 6 // IP = 4 + Port = 2
	if len(peersBin)%peerSize != 0 {
		err := fmt.Errorf("recieved malformed peers")
		return nil, err
	}

	numPeers := len(peersBin) / peerSize

	peers := make([]Peer, numPeers)
	for i := range numPeers {
		offset := i * peerSize
		peers[i].IP = net.IP(peersBin[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16(peersBin[offset+4 : offset+6])
	}

	return peers, nil
}

// belongs to torFile (e.g. torFile.buildTrackerURL(...))
func (torFile *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(torFile.Announce)
	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(torFile.InfoHash[:])},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(torFile.Length)},
	}

	base.RawQuery = params.Encode()

	return base.String(), nil
}

func (p Peer) String() string {
    return fmt.Sprintf("%s:%d", p.IP.String(), p.Port)
}
