package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// A place to store the cancel functions for all of our goroutines.
	var cancelMap []context.CancelFunc

	// Create global context.
	ctx, cancelAll := context.WithCancel(context.Background())

	// Create 3 goroutines.
	for i := 1; i < 4; i++ {
		// Create sub context.
		ctx, cancel := context.WithCancel(ctx)
		// Add cancel func to var outside of goroutine.
		cancelMap = append(cancelMap, cancel)

		go func(i int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println(fmt.Sprintf("stopping worker %d", i))
					return
				default:
					fmt.Println(i)
					time.Sleep(time.Second)
				}
			}
		}(i)
	}

	// Wait 5 seconds, then cancel the goroutine that is printing "1".
	time.Sleep(5 * time.Second)
	if len(cancelMap) > 0 {
		cancelMap[0]()
	}

	// Wait 5 seconds, then cancel the goroutines that are printing "2" and "3".
	time.Sleep(5 * time.Second)
	cancelAll()

	// Wait to allow final print statements to run.
	time.Sleep(2 * time.Second)
}
