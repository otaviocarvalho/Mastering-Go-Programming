package main

import (
	"fmt"
	"time"
)

func main() {
    ticker := time.NewTicker(1 * time.Second)
	done := tickCounter(ticker)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(15 * time.Second)
	fmt.Println("Exiting main ..")
}

func tickCounter(ticker *time.Ticker) chan bool {
    done := make(chan bool)
    go func() {
        i := 0
    Loop:
        for {
            select {
            case t := <-ticker.C:
                i++
                fmt.Println("Count", i, "at", t)
            case <-done:
                fmt.Println("done signal")
                break Loop
            }
        }

        fmt.Println("Exiting tick counter ...")
    }()

    return done
}
