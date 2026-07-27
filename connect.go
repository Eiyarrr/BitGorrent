package bittorrent

import (
	"net"
	"time"
)

type PeerConnection struct {
	Peer Peer
	Conn net.Conn
}

func connect(peer Peer) (*PeerConnection, error) {
	conn, err := net.DialTimeout("tcp", peer.String(), 3*time.Second)
	if err != nil {
		return nil, err
	}

	return &PeerConnection{peer, conn}, nil
}
