package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type BlockInfo struct {
	Height    string
	Time      string
	ChainName string
}

var logs = make(chan interface{})
var blockInfoChan = make(chan BlockInfo)

func init() {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(os.Stderr)

	go func() {
		for msg := range logs {
			msg = strings.TrimRight(strings.TrimLeft(fmt.Sprint(msg), "["), "]")
			log.Println("CMD | ", msg)
		}
	}()
}
func l(v ...any) {

	logs <- v
}
