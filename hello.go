package main

import (
	"fmt"
	"strconv"
	"time"
)

func say(s string, i int) {
	t := strconv.Itoa(i)
	fmt.Println(s + t)
}
func main() {
	for i := 0; i < 5; i++ {
		go say("Hello from goroutine ", i)
		time.Sleep(1 * time.Second)
	}

}
