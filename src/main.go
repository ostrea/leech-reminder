package main

import "github.com/gen2brain/beeep"

func main() {
	err := beeep.Notify("Hello", "World", "")
	if err != nil {
		panic(err)
	}
}
