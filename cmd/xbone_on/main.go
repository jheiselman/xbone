package main

import (
	"flag"
	"fmt"

	"github.com/jheiselman/xbone"
)

func main() {
	var xboxIP = flag.String("hostname", "XboxOne", "Hostname or IP address of Xbox")
	var xboxLiveID = flag.String("liveid", "FD00A44B102D3341", "Xbox Live ID")
	flag.Parse()

	xbox := &xbone.Xbox{IP: *xboxIP, LiveID: *xboxLiveID, Debug: true}
	if err := xbox.TurnOn(); err != nil {
		fmt.Printf("Failed to turn on Xbox: %v\n", err)
	} else {
		fmt.Println("Xbox sent command to turn on")
	}
}
