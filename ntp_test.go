package ntp

import (
	"log"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	remoteTime, err := Time("time.google.com")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("remote time: " + remoteTime.UTC().Format(time.DateTime))
	log.Println("local time: " + time.Now().UTC().Format(time.DateTime))
}
