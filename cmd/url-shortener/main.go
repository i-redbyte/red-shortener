package main

import (
	"fmt"
	"red-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
