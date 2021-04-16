package main

import "ChattingRoom/library"

func main() {
	go library.DistributeMsg()
	library.HandleRequests()
}