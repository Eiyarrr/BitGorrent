package main

import (
	"fmt"

	"github.com/Eiyarrr/BitGorrent/client"
)

func main() {
	torrentPath := "testTorrents/debian-edu-13.6.0-amd64-netinst.iso.torrent"
	c, err := client.New(torrentPath)
	if err != nil {
		fmt.Println(err)
	}

	c.Run()
}
