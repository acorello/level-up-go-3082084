package main

import (
	"flag"
	"fmt"
	"sync"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeatWG(n int, message string) {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(n)
	for i := n; i > 0; i-- {
		go func() {
			fmt.Println(message)
			defer wg.Done()
		}()
	}
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	for _, m := range messages {
		fmt.Println(m)
		repeatWG(int(*factor), m)
	}
}
