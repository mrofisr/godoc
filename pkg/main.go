package pkg

import "github.com/docker/docker/client"

type Client struct {
	cli *client.Client
}

func Run() *Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return &Client{cli}
}
