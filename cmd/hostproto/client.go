package hostproto

import (
	"context"
	"log"
	"time"
)

// Timeout between greetings
const Timeout = 1 * time.Second

// Greet using the given client
func CallHost(c HostClient, log *log.Logger, dest string, message string) error {
	log.Printf("Calling %s CallHost %s", dest, message)
	r, err := c.CallHost(context.Background(), &HostRequest{Message: message})
	if err != nil {
		log.Printf("could not call CallHost: %v", err)
		return err
	}
	log.Printf("Response received from %s: %s", dest, r.Message)

	return nil
}
