package promise

type any interface{}

type Callback func(x any)
type Func func(x any) any

type Promise struct {
	status chan string
	value  chan any
	reason chan any
}

func Create() *Promise {
	p := &Promise{
		status: make(chan string, 1),
		value:  make(chan any, 1),
		reason: make(chan any, 1),
	}

	p.status <- "pending"
	p.value <- nil
	p.reason <- nil

	return p
}

func (p *Promise) Resolve(value any) {
	p.status <- "fulfilled"
	p.value <- value
	p.reason <- nil
}

func (p *Promise) Reject(reason any) {
	p.status <- "reject"
	p.reason <- reason
	p.value <- nil
}

func (p *Promise) Wait() (string, any, any) {
	<-p.status
    <-p.value
	<-p.reason

	s := <-p.status
    v := <-p.value
	r := <-p.reason

	return s, v, r
}

func New(f func(resolve Callback, reject Callback)) *Promise {
	p := Create()

	go f(p.Resolve, p.Reject)

	return p
}

func (p *Promise) Then(resolve Func) *Promise {
	promise := Create()

	go func() {
		status, value, reason := p.Wait()

		if status != "fulfilled" {
			promise.Reject(reason)
		} else {
			result := resolve(value)
			promise.Resolve(result)
		}
	}()

	return promise
}

func (p *Promise) Catch(reject Func) *Promise {
	promise := Create()

	go func() {
		status, value, reason := p.Wait()

		if status != "reject" {
			promise.Resolve(value)
		} else {
			result := reject(reason)
			promise.Reject(result)
		}
	}()

	return promise
}

func (p *Promise) Finally(finally func()) *Promise {
	promise := Create()

	go func() {
		status, value, reason := p.Wait()
		finally()

		if status == "resolve" {
			promise.Resolve(value)
		} else if status == "reject" {
			promise.Reject(reason)
		}
	}()

	return promise
}
