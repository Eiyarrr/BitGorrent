package client

import "fmt"

type Client struct{}

func New() Client {
	fmt.Println("New client created!")

	return Client{}
}

func (c *Client) Run() {
	fmt.Println("Running client!")
}
