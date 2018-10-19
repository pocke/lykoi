package main

import (
	"github.com/pocke/lykoi"
)

func main() {
	err := lykoi.Init()
	if err != nil {
		panic(err)
	}
	defer lykoi.Exit(0)

	select {}
}
