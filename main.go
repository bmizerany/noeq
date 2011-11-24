package main

import (
	"flag"
	"fmt"
	"github.com/bmizerany/noeq.go"
	"log"
)

// Flags
var (
	addr = flag.String("a", "127.0.0.1:4444", "the noeq address")
	c    = flag.Int("c", 1, "the number of ids to generate (1-255)")
)

func main() {
	flag.Parse()

	ne, err := noeq.New(*addr)
	if err != nil {
		log.Fatal(err)
	}

	ids, err := ne.Gen(uint8(*c))
	if err != nil {
		log.Fatal(err)
	}

	for _, id := range ids {
		fmt.Println(id)
	}
}
