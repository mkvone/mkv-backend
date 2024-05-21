package main

import (
	"github.com/mkvone/mkv-backend/cmd"
	"github.com/mkvone/mkv-backend/cmd/config"
)

func main() {
	cfg := config.LoadConfig("")
	// utils.PrettyPrint(cfg)
	cmd.Run(cfg)
}
