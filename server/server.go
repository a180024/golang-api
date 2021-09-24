package server

import (
	"github.com/a180024/nft_api/config"
)

func Init() {
	r := NewRouter()
	c := config.GetConfig()
	r.Run(c.GetString("port"))
}
