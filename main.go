package main

import (
	"QoinTechnicalTest-Golang/games"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	games.PlayDice(5, 4)
}
