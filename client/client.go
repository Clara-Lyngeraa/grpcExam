package main

import (
	"flag"
)

type Client struct {
	id         int32
	portNumber int32
}

var (
	clientPort = flag.Int("cPort", 0, "Client port number")
)
