package util

import (
	"context"
	"fmt"
	"log"

	goApiAbrha "github.com/abrhacom/go-api-abrha"
)

func ExampleWaitForActive() {
	// Create a goApiAbrha client.
	client := goApiAbrha.NewFromToken("dop_v1_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Create a Vm.
	vmRoot, resp, err := client.Vms.Create(context.Background(), &goApiAbrha.VmCreateRequest{
		Name:   "test-vm",
		Region: "nyc3",
		Size:   "s-1vcpu-1gb",
		Image: goApiAbrha.VmCreateImage{
			Slug: "ubuntu-20-04-x64",
		},
	})
	if err != nil {
		log.Fatalf("failed to create vm: %v\n", err)
	}

	// Find the Vm create action, then wait for it to complete.
	for _, action := range resp.Links.Actions {
		if action.Rel == "create" {
			// Block until the action is complete.
			if err := WaitForActive(context.Background(), client, action.HREF); err != nil {
				log.Fatalf("error waiting for vm to become active: %v\n", err)
			}
		}
	}

	fmt.Println(vmRoot.Vm.Name)
}
