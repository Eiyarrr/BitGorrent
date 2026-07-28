package bittorrent

import "io"

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

func Read(reader io.Reader) (*Handshake, error) {
	buf := make([]byte, 68)

	// read full stream into buf
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}

	h := &Handshake{}

	// ProtocolStr
	protoStrLen := int(buf[0])
	h.ProtocolStr = string(buf[1 : 1+protoStrLen])

	curr := 1 + protoStrLen

	// skip 8 reserved bytes
	curr += 8

	// InfoHash
	copy(h.InfoHash[:], buf[curr:curr+20])
	curr += 20

	// PeerID
	copy(h.PeerID[:], buf[curr:curr+20])

	return h, nil
}
