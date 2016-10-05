package futures

type Deferred struct {
	DeferredValue *DeferredValue
}

func (promise *Deferred) init() {
	if promise.DeferredValue == nil {
		promise.DeferredValue = &DeferredValue{}
	}
}

func (promise *Deferred) Resolve() {
	promise.init()
	promise.DeferredValue.PResolve(nil)
}

func (promise *Deferred) Reject(err error) {
	promise.init()
	promise.DeferredValue.PReject(err)
}

func (promise *Deferred) Errors() []error {
	promise.init()
	return promise.DeferredValue.PErrors()
}

func (promise *Deferred) Then(resolve func(), reject func(error)) *Deferred {
	promise.init()
	promise.DeferredValue.PThen(func(value interface{}) {
		resolve()
	}, reject)
	return promise
}

func (promise *Deferred) PThen(result func(interface{}), reject func(error)) Promise {
	promise.init()
	promise.DeferredValue.PThen(result, reject)
	return promise
}

func (promise *Deferred) PErrors() []error {
	promise.init()
	return promise.DeferredValue.PErrors()
}

func (promise *Deferred) PResolve(result interface{}) {
	promise.init()
	promise.DeferredValue.PResolve(result)
}

func (promise *Deferred) PReject(err error) {
	promise.init()
	promise.DeferredValue.PReject(err)
}
