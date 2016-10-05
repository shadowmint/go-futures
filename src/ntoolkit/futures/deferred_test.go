package futures_test

import (
	"ntoolkit/assert"
	"ntoolkit/futures"
	"testing"
)

func TestDeferredIsPromise(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		var _ futures.Promise = &futures.Deferred{}
	})
}

func TestResolveDeferred(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		resolved := false
		DeferredValue := &futures.Deferred{}

		DeferredValue.Then(func() {
			resolved = true
		}, func(_ error) {
			T.Unreachable()
		})
		DeferredValue.Resolve()

		T.Assert(resolved)
	})
}

func TestResolveDeferredLateEvent(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		resolved := false
		DeferredValue := &futures.Deferred{}

		DeferredValue.Resolve()
		DeferredValue.Then(func() {
			resolved = true
		}, func(_ error) {
			T.Unreachable()
		})

		T.Assert(resolved)
	})
}

func TestChaining(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		resolved := 0
		task := (&futures.Deferred{}).Then(func() {
			resolved++
		}, func(_ error) {
			T.Unreachable()
		}).Then(func() {
			resolved++
		}, func(_ error) {
			T.Unreachable()
		})

		task.Resolve()
		T.Assert(resolved == 2)
	})
}
