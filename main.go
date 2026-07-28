package main

import (
	"fmt"

	"github.com/Eiyarrr/BitGorrent/client"
	"github.com/Eiyarrr/BitGorrent/torrent"
)

func main() {
	fmt.Println("Hello, World!")
	c := client.New()
	c.Run()

	tor, err := torrent.Load("testTorrents/debian-edu-13.6.0-amd64-netinst.iso.torrent")
	if err != nil {
		fmt.Println("Error loading torrent")
	}

	fmt.Println("heres the torrent")
	fmt.Println(tor)
}
