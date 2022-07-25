package main

import (
	"fmt"
	"leech-reminder/src/anki"
	"log"
	"math/rand"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	leeches, err := anki.GetLeeches()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().Unix())
	var usedThisSession = make(map[string]struct{})
	for {
		var randomLeech = leeches[rand.Intn(len(leeches))]
		if _, ok := usedThisSession[randomLeech.Expression]; ok {
			continue
		}

		err := beeep.Notify("Leeeeech!", randomLeech.Expression, "")
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(5 * time.Second)

		answer := fmt.Sprintf("%s\n%s", randomLeech.Reading, randomLeech.Definition)
		err = beeep.Notify("Answer", answer, "")
		if err != nil {
			log.Fatal(err)
		}

		usedThisSession[randomLeech.Expression] = struct{}{}
		time.Sleep(10 * time.Minute)
	}
}
