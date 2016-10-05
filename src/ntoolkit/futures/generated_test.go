package futures_test

//go:generate go run gen/gen.go -packageName futures_test -typeImport testing -typeName T   -typeValue *testing.T -output generated_t_test.go
//go:generate go run gen/gen.go -packageName futures_test                     -typeName Int -typeValue int        -output generated_int_test.go

import (
	"ntoolkit/assert"
	"ntoolkit/futures"
	"testing"
)

func TestDeferredValueInt(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		instance := &DeferredValueInt{}
		instance.Then(func(v int) {
			T.Assert(v == 10)
		}, func(e error) {
			T.Unreachable()
		})
		instance.Resolve(10)
	})
}

func TestDeferredValueIntIsPromise(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		var _ futures.Promise = &DeferredValueInt{}
	})
}

func TestDeferredValueT(tests *testing.T) {
	assert.Test(tests, func(T *assert.T) {
		resolved := false
		instance := &DeferredValueT{}
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

func TestDeferredValueTIsPromise(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		var _ futures.Promise = &DeferredValueT{}
	})
}
