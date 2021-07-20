package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sgarcez/gmapsign"
)

const sigKeyEnvVarName = "GMAPSIGN_KEY"

var sigKeyFlag = flag.String("key", "", "Signing Key (overrides GMAPS_SIGNER_KEY)")

func main() {
	flag.Parse()

	signKey, err := loadSigningKey()
	if err != nil {
		log.Fatal(err)
	}

	if err := gmapsign.Pipeline(os.Stdin, os.Stdout, signKey); err != nil {
		log.Fatal(err)
	}
}

func loadSigningKey() ([]byte, error) {
	raw := *sigKeyFlag
	if raw == "" {
		raw = os.Getenv(sigKeyEnvVarName)
	}

	if raw == "" {
		return nil, fmt.Errorf("signing key expected via -key or %v", sigKeyEnvVarName)
	}

	key, err := gmapsign.DecodeSigningKey(raw)
	if err != nil {
		return nil, fmt.Errorf("decoding signing key: %w", err)
	}

	return key, nil
}
