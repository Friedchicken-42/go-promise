
# Go-Promise

A (*useless*) promise system for handling async tasks in Go


## Installation


```sh
  go get github.com/Friedchicken-42/go-promise
```
## Usage

```go
import (
    promise "github.com/Friedchicken-42/go-promise"
)
```

Then create a Promise
```go
p := promise.New(func(resolve, reject promise.Callback) {
    // other code
    resolve(...)
})
```
and wait for the fulfillment
```go
p.Then(func(v any) any {
    //v is the value passed to resolve
})
```
## Functions

- New
- Then
- Catch
- Finally
- All
- AllSettled
- Any
- Race

### Chain promises
```go
promise.New(func(res, rej promise.Callback) {
    ...
}).Then(func(v any) any {
    ...
}).Catch(func(v any) any {
    ...
})
```

### All
Wait unti all promises are resolved
```go
promise1 = promise.New(func(res, rej promise.Callback) {
    time.Sleep(10 * time.Millisecond)
    resolve(10)
})

promise2 = promise.New(func(res, rej promise.Callback) {
    time.Sleep(20 * time.Millisecond)
    resolve(20) 
})

All([]*promise.Promise{promise1, promise2}).Then(func(v any) any {
    ...
})
```

### AllSettled
Wait until all promises are resolved or rejected

### Any
Wait until the first promise resolve or rehect

### Race
Wait until the first promise resolve
## TODO

- [ ] Type for `[]*Promise`
- [ ] Better go1.18 typing support

