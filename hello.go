package main

import (
	"github.com/gen2brain/beeep";
)


func main() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	 if err != nil {
		panic(err)
	}
}