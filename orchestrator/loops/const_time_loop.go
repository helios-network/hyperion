package loops

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
)

// ErrGracefulStop is a special error, if returned from within loop function, will stop that loop without
// returning any error
var ErrGracefulStop = errors.New("stop")

// Loop runs a function in the loop with a consistent interval. If execution takes longer,
// the waiting time between iteration decreases. A single iteration has a deadline and cannot run longer
// than interval itself. There is a protection from panic which could crash adjacent loops.
func RunLoop(ctx context.Context, interval time.Duration, fn func() error) (err error) {
	defer panicRecover(&err)

	delayTimer := time.NewTimer(0)
	for {
		select {
		case <-delayTimer.C:
			var start = time.Now()
			if fnErr := fn(); fnErr != nil {
				log.WithError(fnErr).Errorln("loop function returned an error")
				// if fnErr == ErrGracefulStop {
				// 	return nil
				// }

				// return fnErr
			}

			if elapsed := time.Since(start); elapsed >= interval {
				// in case of an overlap, use just gstinterval
				delayTimer.Reset(interval)
			} else {
				delayTimer.Reset(interval - elapsed)
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func panicRecover(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = e

			log.WithError(e).Errorln("loop panicked with an error")
			log.Debugln(string(debug.Stack()))
			return
		}

		*err = errors.Errorf("loop panic: %v", r)
		log.Errorln(*err)
	}
}

func RetryFunction[T any](ctx context.Context, callback func() (T, error), delay time.Duration) (p T, err error) {
	var zero T

	// Set up panic recovery
	defer func() {
		if r := recover(); r != nil {
			p = zero
			err = fmt.Errorf("panic in callback: %v", r)
			log.Errorf("Callback function panicked: %v", r)
			// Optional: include stack trace
			debug.PrintStack()
		}
	}()

	for {
		// Check if context is cancelled before each attempt
		select {
		case <-ctx.Done():
			return zero, ctx.Err()
		default:
			// Continue with retry
		}

		// Execute callback with panic protection
		var result T
		var callErr error

		func() {
			defer func() {
				if r := recover(); r != nil {
					callErr = fmt.Errorf("panic in callback: %v", r)
					log.Errorf("Callback function panicked: %v", r)
				}
			}()
			result, callErr = callback()
		}()

		if callErr == nil {
			return result, nil
		}

		log.Warningf("Attempt failed: %v. Retrying in %v...", callErr, delay)

		// Use timer with context to allow for cancellation during sleep
		timer := time.NewTimer(delay)
		select {
		case <-timer.C:
			// Continue to next iteration
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return zero, ctx.Err()
		}
	}
}

func RetryUntilSuccess[T any, P any](callback func(params P) (T, error), params P, delay time.Duration) (T, error) {
	var attempt int

	for {
		attempt++
		result, err := callback(params)
		if err == nil {
			return result, nil
		}

		log.Warningf("Attempt %d failed: %v. Retrying in %v...\n", attempt, err, delay)
		time.Sleep(delay)
	}
}
