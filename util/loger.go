package util

import (
	"github.com/freeddser/ddser/config"

	"log"
)

func Loger(info interface{}) {
	if config.MustGetString("server.Run_Mode") == "debug" {
		log.Println(info)
	}
}
