package torrent

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Eiyarrr/Bencode-Parser"
)

func Load(pathOrURL string) (*Torrent, error) {
	// crazy permissive, local files are also *url.URL
	u, err := url.Parse(pathOrURL)
	if err != nil {
		return nil, err
	}

	// check if online
	if u.Scheme == "http" || u.Scheme == "https" {
		return loadURL(u)
	}

	return loadPath(pathOrURL)
}

func loadURL(u *url.URL) (*Torrent, error) {
	// resp.Body is the io.Reader() for the response data,
	// containing the contents of the .torrent file
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	return parse(resp.Body)
}

func loadPath(path string) (*Torrent, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return parse(file)
}

func parse(reader io.Reader) (*Torrent, error) {
	decoded, err := bencode.Decode(reader)
	if err != nil {
		return nil, err
	}

	dict, ok := decoded.(map[string]any)
	if !ok {
		return nil, fmt.Errorf(".torrent file must contain a dictionary")
	}

	announce, ok := dict["announce"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid announce field")
	}

	info, ok := dict["info"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("missing or invalid info field")
	}

	torrent := &Torrent{
		Announce: announce,
		Info:     info,
	}

	return torrent, nil
}
