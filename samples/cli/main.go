package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var val int
	var err error

	if len(os.Args) < 2 {
		val = rand.Intn(4) // 0-3
	} else {
		val, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println("invalid arguments")
			os.Exit(1)
		}
	}

	log.Println(Do(val))
}

func Do(val int) int {
	return val
}
