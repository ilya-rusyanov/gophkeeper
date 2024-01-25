package main

import (
	"github.com/ilya-rusyanov/gophkeeper/internal/server/config"
)

func main() {
	config := config.New()
	config.MustParse()
}
