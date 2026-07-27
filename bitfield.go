package bittorrent

// bitfields represent what pieces peers have
type Bitfield []byte

func (bf Bitfield) HasPiece(index int) bool {
	byteIndex := index/8
	offset := byteIndex % 8

	// the bitfield in the byte, at the bit, isolate it, return bit != 0
	return bf[byteIndex]>>(7-offset)&1 != 0
}

func (bf Bitfield) SetPiece(index int) {

}
