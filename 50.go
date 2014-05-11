package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
    z:=complex128(1)
    fmt.Println(z)
    for i:=0; i<28; i++ {
    	z = z - (cmplx.Pow(z,3)-x) / (complex128(3)*cmplx.Pow(z,3))
    }
    
    return z
}

func main() {
    fmt.Println(Cbrt(2))
}