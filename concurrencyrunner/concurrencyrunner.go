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
		mu      sync.Mutex
		results []Result
		err     error
	)
	g, _ := errgroup.WithContext(context.Background())

	for _, fn := range functions {
		fn := fn

		g.Go(func() error {
			defer func() {
				if r := recover(); r != nil {
					var (
						ok bool
					)

					err, ok = r.(error)

					if !ok {
						results = append(results, Result{
							Result: nil,
							Error:  fmt.Errorf("panic error: %v", r)})
					}
				}
			}()

			result, err := fn()

			mu.Lock()
			defer mu.Unlock()

			if err == nil {
				results = append(results, Result{
					Result: result,
					Error:  nil})
			}

			return err
		})
	}

	if err = g.Wait(); err != nil {
		results = append(results, Result{
			Result: nil,
			Error:  err})
	}

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
