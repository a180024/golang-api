package main

import (
	"github.com/a180024/nft_api/db"
	"github.com/a180024/nft_api/server"
)

func main() {
	server.Init()
	db.Init()
}
