package bittorrent

import "encoding/binary"

type messageID uint8

const (
	MsgChoke         messageID = 0
	MsgUnchoke       messageID = 1
	MsgInterested    messageID = 2
	MsgNotInterested messageID = 3
	MsgHave          messageID = 4
	MsgBitfield      messageID = 5
	MsgRequest       messageID = 6
	MsgPiece         messageID = 7
	MsgCancel        messageID = 8
)

type message struct {
	ID      messageID
	Payload []byte
}

// Serializes into form: <length prefix><message ID><payload>
func (m *message) Serialize() []byte {
	// interprets 'nil' as a keep alive message
	if m == nil {
		return make([]byte, 4)
	}

	length := uint32(len(m.Payload)+1) // +1 for ID
	buf := make([]byte, 4+length)

	binary.BigEndian.PutUint32(buf[0:4], length)
	// set ID
	buf[4] = byte(m.ID)

	// set Payload
	copy(buf[5:], m.Payload)

	return buf
}
