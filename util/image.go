package util

import (
	"context"
	"fmt"
	"time"

	goApiAbrha "github.com/abrhacom/go-api-abrha"
)

const (
	// availableFailure is the amount of times we can fail before deciding
	// the check for available is a total failure. This can help account
	// for servers randomly not answering.
	availableFailure = 3
)

// WaitForAvailable waits for an image to become available
func WaitForAvailable(ctx context.Context, client *goApiAbrha.Client, monitorURI string) error {
	if len(monitorURI) == 0 {
		return fmt.Errorf("create had no monitor uri")
	}

	completed := false
	failCount := 0
	for !completed {
		action, _, err := client.ImageActions.GetByURI(ctx, monitorURI)

		if err != nil {
			select {
			case <-ctx.Done():
				return err
			default:
			}
			if failCount <= availableFailure {
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
