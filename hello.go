package main

import (
	"fmt"
	"log"
	"os"

	"github.com/keixdata/message-store-go/eventstore"
)

func main() {
	ev, err := eventstore.WithPgConnString(os.Getenv("PG_CONN"))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println(ev)
	data := map[string]string{}
	metadata := map[string]string{}

	msg := eventstore.NewMessage("Test", "Testing", data, metadata)
	_, err = ev.Write(msg)
	log.Fatal(err)
}
