package promise

import (
	"testing"
	"time"
)

func Check(t *testing.T, received, expected int) {
	if received != expected {
		t.Errorf("got %d, wanted %d", expected, received)
	}
}

func TestPromiseThen(t *testing.T) {
	promise := New(func(resolve, _ Callback) {
		time.Sleep(100 * time.Millisecond)
		resolve(100)
	})

	promise.Then(func(value any) any {
		result := value.(int)
		Check(t, result, 100)
		return nil
	})

	time.Sleep(200 * time.Millisecond)
}

func TestPromiseCatch(t *testing.T) {
	promise := New(func(_, reject Callback) {
		time.Sleep(100 * time.Millisecond)
		reject(100)
	})

	promise.Catch(func(value any) any {
		result := value.(int)
		Check(t, result, 100)
		return nil
	})

	time.Sleep(200 * time.Millisecond)
}

func TestPromiseReturn(t *testing.T) {
	promise := New(func(resolve, _ Callback) {
		time.Sleep(100 * time.Millisecond)
		resolve(100)
	})

	promise.Then(func(value any) any {
		return nil
	}).Catch(func(value any) any {
		t.Errorf("unexpected rejection")
		return nil
	})

	time.Sleep(200 * time.Millisecond)
}

func TestPromiseFinally(t *testing.T) {
	promise := New(func(resolve, _ Callback) {
		time.Sleep(100 * time.Millisecond)
		resolve(100)
	})

	called := false

	promise.Catch(func(value any) any {
		t.Errorf("unexpected rejection")
		return nil
	}).Finally(func() {
		called = true
	})

	time.Sleep(200 * time.Millisecond)
	if !called {
		t.Errorf("failed not called")
	}
}

func TestPromiseMultiple(t *testing.T) {
	promise := New(func(resolve, _ Callback) {
		time.Sleep(100 * time.Millisecond)
		resolve(100)
	})

	promise.Then(func(value any) any {
		Check(t, value.(int), 100)
		return value.(int) + 20
	}).Then(func(value any) any {
		Check(t, value.(int), 120)
		return value.(int) + 3
	}).Then(func(value any) any {
		Check(t, value.(int), 123)
		return nil
	}).Catch(func(value any) any {
		t.Errorf("unexpected rejection")
		return nil
	})

	time.Sleep(200 * time.Millisecond)
}
