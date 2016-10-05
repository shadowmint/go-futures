package futures_test

//go:generate go run gen/gen.go -packageName futures_test -typeImport testing -typeName T   -typeValue *testing.T -output generated_t_test.go
//go:generate go run gen/gen.go -packageName futures_test                     -typeName Int -typeValue int        -output generated_int_test.go

import (
	"ntoolkit/assert"
	"ntoolkit/futures"
	"testing"
)

func TestDeferredInt(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		instance := &DeferredInt{}
		instance.Then(func(v int) {
			T.Assert(v == 10)
		}, func(e error) {
			T.Unreachable()
		})
		instance.Resolve(10)
	})
}

func TestDeferredIntIsPromise(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		var _ futures.Promise = &DeferredInt{}
	})
}

func TestDeferredT(tests *testing.T) {
	assert.Test(tests, func(T *assert.T) {
		resolved := false
		instance := &DeferredT{}
		instance.Then(func(v *testing.T) {
			T.Assert(v == tests)
			resolved = true
		}, func(e error) {
			T.Unreachable()
		})
		instance.Resolve(tests)
		T.Assert(resolved)
	})
}

func TestDeferredTIsPromise(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		var _ futures.Promise = &DeferredT{}
	})
}
