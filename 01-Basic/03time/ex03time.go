package main

import (
   "fmt"
   "time"
   "github.com/go-zoox/fetch"
)

func main() {
    start := time.Now()
    _, err := fetch.Get("https://www.google.com/")
	if err != nil {
		panic(err)
	}
    fmt.Println(time.Since(start))
}
	