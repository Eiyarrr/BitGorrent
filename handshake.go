package bittorrent

type Handshake struct {
	ProtocolStr string
	InfoHash    [20]byte
	PeerID      [20]byte
}

func (h *Handshake) Serialize() []byte {
	buf := make([]byte, len(h.ProtocolStr)+49)

	buf[0] = byte(len(h.ProtocolStr))
	curr := 1
	curr += copy(buf[curr:], []byte(h.ProtocolStr))
	curr += copy(buf[curr:], make([]byte, 8)) // 8 reserved bytes
	curr += copy(buf[curr:], h.InfoHash[:])
	curr += copy(buf[curr:], h.PeerID[:])

	return buf
}
