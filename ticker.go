package main

import (
	"fmt"
	"time"
)

func tickMsg(msg string, stop chan bool) {
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		if <-stop {
			ticker.Stop()
		}
	}()
	for {
		what := <-ticker.C
		fmt.Println(msg)
		fmt.Println(what)
	}

}
