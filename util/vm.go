package util

import (
	"context"
	"fmt"
	"time"

	goApiAbrha "github.com/abrhacom/go-api-abrha"
)

const (
	// activeFailure is the amount of times we can fail before deciding
	// the check for active is a total failure. This can help account
	// for servers randomly not answering.
	activeFailure = 3
)

// WaitForActive waits for a vm to become active
func WaitForActive(ctx context.Context, client *goApiAbrha.Client, monitorURI string) error {
	if len(monitorURI) == 0 {
		return fmt.Errorf("create had no monitor uri")
	}

	completed := false
	failCount := 0
	for !completed {
		action, _, err := client.VmActions.GetByURI(ctx, monitorURI)

		if err != nil {
			select {
			case <-ctx.Done():
				return err
			default:
			}
			if failCount <= activeFailure {
				failCount++
				continue
			}
			return err
		}

		switch action.Status {
		case goApiAbrha.ActionInProgress:
			select {
			case <-time.After(5 * time.Second):
			case <-ctx.Done():
				return err
			}
		case goApiAbrha.ActionCompleted:
			completed = true
		default:
			return fmt.Errorf("unknown status: [%s]", action.Status)
		}
	}

	return nil
}
