package main

import (
	server "subman/api/server"
	indexer "subman/indexer"
)

func main() {
	go indexer.Start()
	server.Start()
}
