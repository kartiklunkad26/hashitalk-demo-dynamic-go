package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
)

func main() {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithBaseAddress(os.Getenv("VAULT_ADDR")),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("my-token"); err != nil {
		log.Fatal(err)
	}

	// read dynamic username & password
}
