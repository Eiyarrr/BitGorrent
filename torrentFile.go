package bittorrent

import (
	"github.com/Eiyarrr/Bencode-Parser"
	"io"
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

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func Open(reader io.Reader) (*bencodeTorrent, error) {
	benTor := bencodeTorrent{}
	err := bencode.Unmarshal(reader, benTor)
	if err != nil {
		return nil, err
	}

	return &benTor, nil
}

func toTorrentFile(benTor bencodeTorrent) (TorrentFile, error) {
	torrent := TorrentFile{
		Announce:    benTor.Announce,
		//InfoHash
		PieceHashes: splitPieceHashes([]byte(benTor.Info.Pieces)),
		PieceLength: benTor.Info.PieceLength,
		Length:      benTor.Info.Length,
		Name:        benTor.Info.Name,
	}

	return torrent, nil
}

func splitPieceHashes(pieces []byte) [][20]byte {
	var hashes [][20]byte

	for i, _ := range pieces {
		var hash [20]byte
		// copy the next 20 bytes into hash
		copy(hash[:], pieces[i:i+20])
		// put hash onto end of hashes
		hashes = append(hashes, hash)
	}

	return hashes
}
