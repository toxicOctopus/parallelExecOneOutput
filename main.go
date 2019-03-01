package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func runCommand(c chan string, i int) {
	defer wg.Done()
	//defer close(c)
	time.Sleep(time.Duration(i) * time.Second)
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	c <- string(out)
}

func main() {
	routinesCount := 3
	start := time.Now()
	i := 1
	c := make(chan string)
	wg.Add(routinesCount)
	for i <= routinesCount {
		go runCommand(c, i)
		i++
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	for k := range c {
		fmt.Println(k)
	}
	fmt.Printf("Execution time: %v\n", time.Since(start))
}
