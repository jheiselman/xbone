package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jheiselman/xbone"
)

func main() {
	var xboxIP = flag.String("hostname", "XboxOne", "Hostname or IP address of Xbox")
	var xboxLiveID = flag.String("liveid", os.Getenv("XBOX_LIVE_ID"), "Xbox Live ID")
        var debug = flag.Bool("debug", false, "Turn on debug output")
	flag.Parse()

	if *xboxLiveID == "" {
		fmt.Println("You must pass -liveid option or set XBOX_LIVE_ID environment variable")
		return
	}

	xbox := &xbone.Xbox{IP: *xboxIP, LiveID: *xboxLiveID, Debug: *debug}
	if err := xbox.TurnOn(); err != nil {
		fmt.Printf("Failed to turn on Xbox: %v\n", err)
	} else {
		fmt.Println("Xbox sent command to turn on")
	}
}
