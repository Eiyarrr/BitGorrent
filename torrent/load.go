package torrent

import (
	"net/url"
	"os"
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
	return nil, nil
}

func loadPath(path string) (*Torrent, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil
	}

	return parse(data)
}

func parse(data []byte) (*Torrent, error) {
	return nil, nil
}
