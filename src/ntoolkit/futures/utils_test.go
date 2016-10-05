package futures_test

import (
	"errors"
	"ntoolkit/assert"
	"ntoolkit/futures"
	"testing"
)

func TestAll(tests *testing.T) {
	assert.Test(tests, func(T *assert.T) {
		resolved := false
		p1 := &DeferredValueInt{}
		p2 := &DeferredValueT{}

		promise := futures.All(p1, p2)
		promise.Then(func() {
			resolved = true
		}, func(_ error) {
			T.Unreachable()
		})

		p1.Resolve(1)
		p2.Resolve(tests)

		T.Assert(resolved)
	})
}

func TestAllReject(tests *testing.T) {
	assert.Test(tests, func(T *assert.T) {
		rejected := false
		p1 := &DeferredValueInt{}
		p2 := &DeferredValueT{}

		promise := futures.All(p1, p2)
		promise.Then(func() {
			T.Unreachable()
		}, func(_ error) {
			rejected = true
		})

		p1.Resolve(1)
		p2.Reject(errors.New("nope"))

		T.Assert(rejected)
	})
}

func TestAny(tests *testing.T) {
	assert.Test(tests, func(T *assert.T) {
		resolved := false
		p1 := &DeferredValueInt{}
		p2 := &DeferredValueT{}

		promise := futures.Any(p1, p2)
		promise.Then(func() {
			resolved = true
		}, func(_ error) {
			T.Unreachable()
		})

		p1.Reject(errors.New("Nope"))
		p2.Resolve(tests)

		T.Assert(resolved)
	})
}

func TestAnyReject(tests *testing.T) {
	assert.Test(tests, func(T *assert.T) {
		rejected := false
		p1 := &DeferredValueInt{}
		p2 := &DeferredValueT{}

		promise := futures.Any(p1, p2)
		promise.Then(func() {
			T.Unreachable()
		}, func(_ error) {
			rejected = true
		})

		p1.Reject(errors.New("Nope"))
		p2.Reject(errors.New("Nope"))

		T.Assert(rejected)
	})
}
