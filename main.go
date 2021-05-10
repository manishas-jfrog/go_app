package main

import (
	"fmt"
	"math/rand"
	"time"
)

var frogFacts = []string{
	"Turtles are better than frogs in every way.",
        "Frogs are slimy and nobody likes them.",
        "Frogs can jump, so what? so can fleas and they are a parasite.",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	n := rand.Int() % len(frogFacts)
	fmt.Println(frogFacts[n])
}
