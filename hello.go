package main

import (
	"github.com/gen2brain/beeep";
	"github.com/tidwall/buntdb";
)


func main() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}


}