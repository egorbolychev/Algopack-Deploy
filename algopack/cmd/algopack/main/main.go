package main

import (
	"log"

	"algopack/cmd/algopack"
)

func main() {
	if err := algopack.RootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
