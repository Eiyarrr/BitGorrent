package bittorrent

// Max # bytes a request can ask for
const MaxBlockSize = 16384

// Max # unfulfilled requests a client can have
const MaxBacklog = 5

type Torrent struct {
	Peers       []Peer
	PeerID      [20]byte
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}
