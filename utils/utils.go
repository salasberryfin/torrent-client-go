package utils

import (
	"log"
)

// Check is a util function for error validation
func Check(e error, msg string) {
	if e != nil {
		log.Fatal(msg + ": " + e.Error())
	}
}
