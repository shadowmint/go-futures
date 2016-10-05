# go-futures

Provides the scaffolding for generic promises.

The go:generate to generate strongly typed promise objects via:

    //go:generate go run gen/gen.go -packageName futures_test -typeImport testing -typeName T   -typeValue testing.T -output generated_t_test.go
    //go:generate go run gen/gen.go -packageName futures_test -typeName Int -typeValue int -output generated_int_test.go

And using something like:

    go generate ./src/mypackage/...

You can also use the util functions `All` and `Any` to chain together arbitrary
strongly typed promises using the generic `Promise` interface, or implement it manually.

For convenience the basic strongly typed promises for interface{} (`DeferredValue`) and
untyped (`Deferred`) are already provided.

## Usage

    npm install shadowmint/go-futures --save

Then use and chain promises:

    func Tasks() futures.Deferred {
      foo := &DeferredFoo{}
      bar := &DeferredBar{}
      go func() {
        ...
        foo.Resolve(myfoo)
      }()
      go func() {
        ...
        bar.Resolve(mybar)
      }
      return futures.All(foo, bar)
    }

    ...

    promise := Tasks().Then(func() {
      println("Done")
    }, (err error) {
      println(err.Error())
    })

## Troubleshooting...

If promises are not invoking their callbacks, make sure that functions are returning `*DeferredT`, not `DeferredT`.

Copy by value will result in the handlers not being post invoked on non-pointer types.

## Tests

    npm install
    npm run generate
    npm test
