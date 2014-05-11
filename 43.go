package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
    //"fmt"
)

func WordCount(s string) map[string]int {
    tokens:= strings.Fields(s)
    wcmap:=make(map[string]int)
    
    for _,v := range tokens{
        _, ok := wcmap[v]
        if ok == true {
        	wcmap[v]++
        } else {
        	wcmap[v]=1	
        }
    }
    
	return wcmap
}

func main() {
    wc.Test(WordCount)
}