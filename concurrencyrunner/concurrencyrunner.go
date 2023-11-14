package concurrencyrunner

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

func Run[T any](functions []func() (T, error)) ConcurrencyRunner {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		results []Result
	)
	g, _ := errgroup.WithContext(context.Background())

	for _, fn := range functions {
		wg.Add(1)
		fn := fn

		g.Go(func() error {
			defer wg.Done()
			result, err := fn()

			mu.Lock()
			defer mu.Unlock()

			if err == nil {
				results = append(results, Result{
					Result: result,
					Error:  nil})
			} else {
				results = append(results, Result{
					Result: nil,
					Error:  err})
			}

			return err
		})
	}

	wg.Wait()
	_ = g.Wait()

	return ConcurrencyRunner{Results: results}
}

func (crr ConcurrencyRunner) CombineErrors() error {
	var combinedErrors []string

	for _, result := range crr.Results {
		if result.Error != nil {
			combinedErrors = append(combinedErrors, result.Error.Error())
		}
	}

	if len(combinedErrors) > 0 {
		return fmt.Errorf(strings.Join(combinedErrors, "\n"))
	}

	return nil
}
