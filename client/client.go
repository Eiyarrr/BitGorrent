package client

import (
	"fmt"
	"io"
	"os"
)

type Client struct {
	path   string
	reader io.Reader
}

func New(torrentPath string) (Client, error) {
	fmt.Println("New client created!")
	fmt.Println("Path: ", torrentPath)

	reader, err := os.Open(torrentPath)
	if err != nil {
		return Client{}, err
	}

	c := Client{
		path:   torrentPath,
		reader: reader,
	}

	return c, nil
}

func (c *Client) Run() {
	fmt.Println("Running client!")
}
