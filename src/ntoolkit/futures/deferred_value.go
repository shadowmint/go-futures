package futures

import (
	"errors"
	"fmt"
)

const (
	pending  = 0
	resolved = 1
	rejected = 2
)

// Pending action either reject or resolve
type pendingDeferred struct {

	// The accept Deferred
	resolved func(interface{})

	// The reject Deferred
	rejected func(error)
}

// DeferredValue implements Promise, either locally or remotely.
type DeferredValue struct {

	// The state this promise is currently in.
	state int

	// The error if rejected
	err error

	// The result if resolved
	result interface{}

	// Pending actions
	pending []pendingDeferred

	// Errors that happened while processing events
	errors []error

  // If set, silently consume errors to prevent bubbling panics
  DontPanic bool
}

// PErrors returns the set of errors saved
func (promise *DeferredValue) PErrors() []error {
	return promise.errors
}

// PThen adds callbacks to be invoked if the promise is completed
func (promise *DeferredValue) PThen(resolve func(interface{}), reject func(error)) Promise {
	promise.pending = append(promise.pending, pendingDeferred{
		resolved: resolve,
		rejected: reject})
	promise.flush()
	return promise
}

// PResolve resolves, runs any pending actions
func (promise *DeferredValue) PResolve(result interface{}) {
	if promise.state != pending {
		return
	}
	promise.result = result
	promise.state = resolved
	promise.flush()
}

// PReject resolves, runs any pending actions
func (promise *DeferredValue) PReject(err error) {
	if promise.state != pending {
		return
	}
	promise.err = err
	promise.state = rejected
	promise.flush()
}

// Flush any pending actions
func (promise *DeferredValue) flush() {
	if promise.state == pending {
		return
	}
	callbacks := promise.pending[:]
	for i := 0; i < len(callbacks); i++ {
		if err := promise.safeInvoke(callbacks[i]); err != nil {
			promise.errors = append(promise.errors, err)
		}
	}
  if len(promise.errors) > 0 {
    if !promise.DontPanic {
      panic(promise.errors)
    }
  }
	promise.pending = nil
}

// Safely invoke a rejected
func (promise *DeferredValue) safeInvoke(callback pendingDeferred) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New(fmt.Sprintf("%s", r))
			}
		}
	}()
	if promise.state == resolved {
		callback.resolved(promise.result)
	} else if promise.state == rejected {
		callback.rejected(promise.err)
	}
	return nil
}
