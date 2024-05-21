package db

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var logs = make(chan interface{})

func init() {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(os.Stderr)

	go func() {
		for msg := range logs {
			msg = strings.TrimRight(strings.TrimLeft(fmt.Sprint(msg), "["), "]")
			log.Println("mongo | ", msg)
		}
	}()
}
func l(v ...any) {

	logs <- v
}
