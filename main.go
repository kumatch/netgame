package main

import (
	"math/rand"
	"time"

	"github.com/kumatch/netgame/quest"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quest.Run()
}
