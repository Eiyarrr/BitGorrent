package main

import (
	"fmt"
	"github.com/Eiyarrr/BitGorrent/client"
)

func main() {
	fmt.Println("Hello, World!")
	c := client.New()
	c.Run()
}
