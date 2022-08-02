package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}

func Shout(msg string) {
    louder := strings.ToUpper(msg)
    fmt.Println(louder)
}
    
