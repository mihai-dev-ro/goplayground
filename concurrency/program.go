package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// testGoroutines()
	// testChannels()
	// testChannelsSelect()
	testSleep()
}

func testGoroutines() {
	for i := 0; i < 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ": ", i)
		duration := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * duration)
	}
}

func testChannels() {
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Println("Send ENTER:")
	fmt.Scanln(&input) 
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping (%v)", i)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("pong (%v)", i)
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println("Received: ", msg)
		time.Sleep(time.Millisecond * 1000)
	}
}

func testChannelsSelect() {
	c1 := make(chan string, 5)
	c2 := make(chan string, 3)

	go func() {
		for i := 0; ; i++ {
			c1 <- fmt.Sprintf("channel 1 msg (%v)", i)
			time.Sleep(time.Millisecond * 250)
		}
	}()

	go func() {
		for i := 0; ; i++ {
			c2 <- fmt.Sprintf("channel 2 msg (%v)", i)
			time.Sleep(time.Millisecond * 301)
		}
	}()

	go func() {
		for {
			select {
			case msg := <- c1:
				fmt.Println(msg)
			case msg := <- c2:
				fmt.Println(msg)
			case msg := <- time.After(time.Second):
				fmt.Println("timeout at ", msg)
			// default:
			// 	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func testTimeout() {
	var index int
	fnCallback := func(eventTime time.Time) {
		index++
		fmt.Printf("timeout (%v) at %v\n", index, eventTime)
	}

	timeout(time.Millisecond * 1000, fnCallback)

	var input string
	fmt.Scanln(&input)
}

func timeout(duration time.Duration, fn func(time.Time)) {
	go func() {
		for {
			time := <- time.After(duration)
			fn(time)
		}
	}()
}

func testSleep() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Iteration (%v)\n", i+1)
		sleep(time.Millisecond * 1000)
	}
}

func sleep(duration time.Duration) {
	c := make(chan bool, 2)

	go func() {
		var _ = <- time.After(duration) 
		c <- true
	}()

	fmt.Printf("sleep is in freeze mode, wainting for the data to be available on the channel\n")
	var _ = <- c
	fmt.Printf("sleep ended\n")
}
