package futures

// Promise is a generic abstract base for creating strongly typed promises.
// Although this interface can be used with the generic 'DeferredValue' struct,
// see `typed_promise_test.go` for an example of how it should be used.
type Promise interface {

	// Then adds callbacks to be invoked when the promise is complete.
	// Returns the promise when called, for chaining.
	PThen(func(interface{}), func(error)) Promise

	// Errors returns the set of errors the occurred processing the event
	PErrors() []error

	// Resolve this promise
	PResolve(interface{})

	// Reject this promise
	PReject(error)
}
