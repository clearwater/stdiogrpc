package pluginproto

import (
	"context"
	"log"
	"time"
)

// Timeout between greetings
const Timeout = 1 * time.Second

// Greet using the given client
func CallPlugin(c PluginClient, log *log.Logger, dest string, message string) error {
	log.Printf("Calling %s CallPlugin %s", dest, message)
	r, err := c.CallPlugin(context.Background(), &PluginRequest{Message: message})
	if err != nil {
		log.Printf("could not call CallPlugin: %v", err)
		return err
	}
	log.Printf("Response received from %s: %s", dest, r.Message)

	return nil
}
