package main

import "flag"

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8261", "grpc server addr")
	flag.Parse()
}

func main() {
	server := newServer(addr)
	server.Listen()
}
