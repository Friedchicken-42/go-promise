package promise

import (
	"testing"
	"time"
)

func TestAll(t *testing.T) {
    promise1 := New(func(resolve, reject Callback) {
        time.Sleep(50 * time.Millisecond)
        resolve(10)
    })

    promise2 := New(func(resolve, reject Callback) {
        time.Sleep(50 * time.Millisecond)
        resolve(20)
    })

    result := [2]int{}

    All([]*Promise{promise1, promise2}).Then(func(value any) any {
        for i, v := range value.([]any) {
            result[i] = v.(int)
        }
        return nil
    })


    time.Sleep(100 * time.Millisecond)
    if result != [2]int{10, 20} {
        t.Errorf("wrong result: %+v\n", result)
    }
}

func TestAllSettled(t *testing.T) {
    promise1 := New(func(resolve, reject Callback) {
        time.Sleep(50 * time.Millisecond)
        resolve(10)
    })

    promise2 := New(func(resolve, reject Callback) {
        time.Sleep(50 * time.Millisecond)
        reject(20)
    })

    result := [2]int{}

    AllSettled([]*Promise{promise1, promise2}).Then(func(value any) any {
        for i, v := range value.([]any) {
            result[i] = v.(int)
        }
        return nil
    })

    time.Sleep(100 * time.Millisecond)
    if result != [2]int{10, 20} {
        t.Errorf("wrong result: %+v\n", result)
    }
}

func TestAny(t *testing.T) {
    promise1 := New(func(resolve, reject Callback) {
        time.Sleep(60 * time.Millisecond)
        resolve(10)
    })

    promise2 := New(func(resolve, reject Callback) {
        time.Sleep(40 * time.Millisecond)
        resolve(20)
    })

    result := 0
    
    Any([]*Promise{promise1, promise2}).Then(func(value any) any {
        result = value.(int)
        return nil
    })

    time.Sleep(100 * time.Millisecond)

    if result != 20 {
        t.Errorf("wrong result: %+v\n", result)
    }
}
