package main

import (
	"github.com/svcp/server"
)

func main() {
	s := server.NewServer("/", "8383")
	s.Run()
}
