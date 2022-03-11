package promise

func All(promises []*Promise) *Promise {
	promise := Create()

	go func() {
		result := make([]any, len(promises))
		done := make(chan bool, len(promises))

		for i := range promises {
			_i := i
			promises[i].Then(func(value any) any {
				result[_i] = value
				done <- true
				return nil
			}).Catch(func(value any) any {
				promise.Reject(value)
				return nil
			})
		}

		for range promises {
			<-done
		}

		promise.Resolve(result)
	}()

	return promise
}

func AllSettled(promises []*Promise) *Promise {
	promise := Create()

	go func() {
		result := make([]any, len(promises))
		done := make(chan bool, len(promises))

		for i := range promises {
			_i := i
			promises[i].Then(func(value any) any {
				result[_i] = value
				done <- true
				return nil
			}).Catch(func(value any) any {
				result[_i] = value
				done <- true
				return nil
			})
		}

		for range promises {
			<-done
		}

		promise.Resolve(result)
	}()

	return promise
}

func Any(promises []*Promise) *Promise {
	promise := Create()

	go func() {
		errors := make([]any, len(promises))
		e := make(chan any, 1)
		v := make(chan any, 1)
		done := make(chan bool, len(promises))

		for i := range promises {
			_i := i
			promises[_i].Then(func(value any) any {
				done <- true
				v <- value
				return nil
			}).Catch(func(value any) any {
				done <- false
				e <- value
				return nil
			})
		}

		i := 0
		for range promises {
			s := <-done
			if s {
				x := <-v
				promise.Resolve(x)
				return
			} else {
				errors[i] = <-e
				i++
			}
		}

		promise.Reject(errors)
	}()

	return promise
}

func Race(promises []*Promise) *Promise {
	promise := Create()

	go func() {
        status := make(chan bool, 1)
        result := make(chan any, 1)

        for i := range promises {
            _i := i
            promises[_i].Then(func(value any) any {
                status <- true
                result <- value
                return nil
            }).Catch(func(value any) any {
                status <- false
                result <- value
                return nil
            })
        }

        s := <- status
        v := <- result
        if s {
            promise.Resolve(v)
        } else {
            promise.Reject(v)
        }
	}()

	return promise
}
