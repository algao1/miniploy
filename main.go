package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		t := <-ticker.C
		fmt.Println(t)
	}
}
