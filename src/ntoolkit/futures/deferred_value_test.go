package futures_test

import (
	"errors"
	"ntoolkit/assert"
	"ntoolkit/futures"
	"testing"
	"time"
)

func TestCreatePromise(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		var _ futures.Promise = &futures.DeferredValue{}
	})
}

func TestResolveActionInvokedOnResolve(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		resolveCalled := false
		DeferredValue := &futures.DeferredValue{}
		DeferredValue.PThen(func(data interface{}) {
			v, ok := data.(bool)
			T.Assert(ok)
			T.Assert(v == true)
			resolveCalled = true
		}, func(err error) {
			T.Unreachable()
		})
		DeferredValue.PResolve(true)
		T.Assert(resolveCalled)
	})
}

func TestResolvedActionInvokedWhenAlreadyResolved(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		resolveCalled := false
		DeferredValue := &futures.DeferredValue{}
		DeferredValue.PResolve(true)
		DeferredValue.PThen(func(data interface{}) {
			v, ok := data.(bool)
			T.Assert(ok)
			T.Assert(v == true)
			resolveCalled = true
		}, func(err error) {
			T.Unreachable()
		})
		T.Assert(resolveCalled)
	})
}

func TestRejectActionInvokedOnReject(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		rejectCalled := false
		DeferredValue := &futures.DeferredValue{}
		DeferredValue.PThen(func(data interface{}) {
			T.Unreachable()
		}, func(err error) {
			T.Assert(err.Error() == "Failed")
			rejectCalled = true
		})
		DeferredValue.PReject(errors.New("Failed"))
		T.Assert(rejectCalled)
	})
}

func TestRejectActionInvokedWhenAlreadyRejected(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		rejectCalled := false
		DeferredValue := &futures.DeferredValue{}
		DeferredValue.PReject(errors.New("Failed"))
		DeferredValue.PThen(func(data interface{}) {
			T.Unreachable()
		}, func(err error) {
			T.Assert(err.Error() == "Failed")
			rejectCalled = true
		})
		T.Assert(rejectCalled)
	})
}

func TestAsyncPromiseResolution(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		counter := 0
		DeferredValue := &futures.DeferredValue{}

		DeferredValue.PThen(func(data interface{}) {
			v, ok := data.(int)
			T.Assert(ok)
			T.Assert(v == 4950)
			counter++
		}, func(err error) {
			T.Unreachable()
		})

		DeferredValue.PThen(func(data interface{}) {
			v, ok := data.(int)
			T.Assert(ok)
			T.Assert(v == 4950)
			counter++
		}, func(err error) {
			T.Unreachable()
		})

		go func() {
			total := 0
			for i := 0; i < 100; i++ {
				total += i
			}
			DeferredValue.PResolve(total)
		}()

		time.Sleep(time.Millisecond * 100)
		T.Assert(counter == 2)
	})
}

func TestErrorsAreCollected(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		counter := 0
    DeferredValue := &futures.DeferredValue{DontPanic: true}

		DeferredValue.PThen(func(data interface{}) {
			panic(errors.New("Test"))
		}, func(err error) {
			T.Unreachable()
		})

		DeferredValue.PThen(func(data interface{}) {
			v, ok := data.(int)
			T.Assert(ok)
			T.Assert(v == 4950)
			counter++
		}, func(err error) {
			T.Unreachable()
		})

		go func() {
			total := 0
			for i := 0; i < 100; i++ {
				total += i
			}
			DeferredValue.PResolve(total)
		}()

		time.Sleep(time.Millisecond * 100)
		T.Assert(counter == 1)
		T.Assert(len(DeferredValue.PErrors()) == 1)
		T.Assert(DeferredValue.PErrors()[0].Error() == "Test")
	})
}
