package helloworld

import (
	"context"
	"log"
	"time"
)

// Timeout between greetings
const Timeout = 1 * time.Second

// Greet using the given client
func Greet(c GreeterClient, log *log.Logger, dest string, name string) error {
	log.Printf("Calling %s SayHello %s", dest, name)
	r, err := c.SayHello(context.Background(), &HelloRequest{Name: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
		return err
	}
	log.Printf("Response received from %s: %s", dest, r.Message)

	log.Printf("Calling %s SayHelloAgain %s", dest, name)
	r, err = c.SayHelloAgain(context.Background(), &HelloRequest{Name: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
		return err
	}
	log.Printf("Response received from %s: %s", dest, r.Message)
	return nil
}
