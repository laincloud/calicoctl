package calicoclient

import (
	"fmt"
	"os"

	"github.com/projectcalico/libcalico-go/lib/api"
	"github.com/projectcalico/libcalico-go/lib/client"
)

func CreateClient(filename string) (*api.CalicoAPIConfig, *client.Client) {
	cfg, err := client.LoadClientConfig(filename)
	if err != nil {
		fmt.Printf("ERROR: Error loading datastore config: %s", err)
		os.Exit(1)
	}
	c, err := client.New(*cfg)
	if err != nil {
		fmt.Printf("ERROR: Error accessing the Calico datastore: %s", err)
		os.Exit(1)
	}

	return cfg, c
}
