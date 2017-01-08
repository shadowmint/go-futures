package futures

import (
	"errors"
	"fmt"
)

// Returns a Deferred that resolves if all input promises pass.
func All(promises ...Promise) *Deferred {
	DeferredValue := &Deferred{}
	if len(promises) == 0 {
		DeferredValue.Resolve()
		return DeferredValue
	}
	count := 0
	target := len(promises)
	rejected := 0
	errList := make([]error, 0)
	resolver := func(success bool, err error) {
		if !success {
			rejected++
			errList = append(errList, err)
		}
		count++
		if count == target {
			if rejected > 0 {
				DeferredValue.Reject(errors.New(fmt.Sprintf("%d/%d Deferreds failed: %s", rejected, count, errList)))
			} else {
				DeferredValue.Resolve()
			}
		}
	}
	for _, promise := range promises {
		promise.PThen(func(_ interface{}) {
			resolver(true, nil)
		}, func(err error) {
			resolver(false, err)
		})
	}
	return DeferredValue
}

// Returns a Deferred that resolves if any input promises pass.
func Any(promises ...Promise) *Deferred {
	DeferredValue := &Deferred{}
	if len(promises) == 0 {
		DeferredValue.Resolve()
		return DeferredValue
	}
	count := 0
	target := len(promises)
	resolved := false
	resolver := func(success bool) {
		if success {
			resolved = true
		}
		count++
		if count == target {
			if resolved {
				DeferredValue.Resolve()
			} else {
				DeferredValue.Reject(errors.New("No Deferreds resolved"))
			}
		}
	}
	for _, promise := range promises {
		promise.PThen(func(_ interface{}) {
			resolver(true)
		}, func(_ error) {
			resolver(false)
		})
	}
	return DeferredValue
}
