package main

import "log"

func main() {
	if err := runApp(); err != nil {
		log.Fatalf("runApp: %s", err.Error())
	}
}
