package bittorrent

type Handshake struct {
	ProtocolStr string
	InfoHash [20]byte
	PeerID [20]byte
}
