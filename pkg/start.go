package pkg

import "github.com/docker/docker/client"

func DockerStart() *Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return &Client{cli}
}
