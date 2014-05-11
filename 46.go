package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    fib0:=0
    fib1:=0
    return func() int{
        if fib0==0 && fib1==0 {
            fib0=0
            fib1=1
            return fib1
        } else{ 
        	p := fib1+fib0
        	fib0 = fib1
        	fib1 = p
        	return fib1
        }
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}