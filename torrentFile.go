package bittorrent

import (
	"io"
	"github.com/Eiyarrr/Bencode-Parser"
)

type bencodeInfo struct {
    Pieces      string `bencode:"pieces"`
    PieceLength int    `bencode:"piece length"`
    Length      int    `bencode:"length"`
    Name        string `bencode:"name"`
}

type bencodeTorrent struct {
    Announce string      `bencode:"announce"`
    Info     bencodeInfo `bencode:"info"`
}

func Open(reader io.Reader) (*bencodeTorrent, error) {
	benTor := bencodeTorrent{}
	err := bencode.Unmarshal(reader, benTor)
	if err != nil {
		return nil, err
	}

	return &benTor, nil
}
