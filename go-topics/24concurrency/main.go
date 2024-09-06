package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// What is concurrency
	/*
		Concurrency is the ability to perform multiple task at the same time.
		Concurrency help us achieve concurrent execution which means :-
		two task can execute at the same time.

		Sequential Execution is the opposite of the  Concurrent execution.
		where second task is only executed when first task ends.

		A goroutine is a lightweight thread managed by the Go runtime.
		It's a function or method that is capable of running concurrently with other functions or methods.

		The go keyword is used to start a new goroutine.
		Goroutines are a fundamental part of Go's concurrency model and they are used to perform tasks concurrently

		when a goroutine is spawned with a go keyword and it runs a function in parallel we cannot fetch return value from the function

	*/

	/*
		difference between between parallelism and concurrency
		parallelism refers two task executing independently in different isolated environment independent of each other where order does not matter

		concurrency refers to the
	*/

	/*
	 creating a go routine which prints message mail received after the mail sent
	*/

	/*
		the objective of the below two functions is to
		display mail received by the recipient msg
		after the mail send to the recipient msg
	*/
	mailReceived := func(recipientName string) {
		fmt.Printf("mail received by %s \n", recipientName)
	}

	mailSent := func(recipientName string) {
		fmt.Printf("mail sent to %s \n", recipientName)
	}

	fmt.Println("Explanation 1")
	go mailReceived("Pascal")
	mailSent("Pascal")

	/*
		when we run mailReceived() as a separate goroutine
		when we run mailSent() in the main goroutine
		output: "mail sent to Pascal"
		we will never see the mailReceived() function running in separate goroutine
		reason : as soon as the main function goroutine ends all the goroutine running in the background end
		it does not wait for other running goroutines
	*/

	fmt.Println("Explanation 2")

	go mailReceived("Travis")
	go mailSent("Travis")
	// main function goroutine will go to sleep giving all the other go routine to finish there execution
	time.Sleep(time.Millisecond * 40)

	/*
		when we run mailReceived() in a separate go routine
		when we run mailSent() in a separate go routine
		output:""
		no output as both the function are in separate goroutine
		as soon as the main function goroutine ends all running goroutines gets terminated
		main function() go routine does not wait for the completion of other goroutines
		how to fix it ?
	*/

	fmt.Println("Explanation 3")
	/*
		in order to make main function goroutine wait for other goroutine to complete there execution
		we implement a blocking call at the end of the main function this can be done
		by implementing
		1. a fmt.Scanln() # that waits for the user input to exit / terminate the main function goroutine
		we don't use this approach
		2. making the main function sleep for some time time.Sleep(time.Millisecond * 40)
		3. waitGroup() has mainly three methods wg.Add(n) where n is the number of the goroutines that you have decided to run
		4. wg.Done() decrements the waitGroup counter(n) by 1
		5. wg.Wait() will halt the further execution until WaitGroup counter(n) reaches 0
	*/

	/*
		waitGroup is use to halts the execution of the main goroutine
		and helps in execution of all the other go routine running in the background.
		waitGroup has two functions
	*/

	// Package sync provides basic synchronization primitives such as mutual exclusion locks or WaitGroup types
	var wg sync.WaitGroup
	wg.Add(2) // add counter tells the number of goroutine it has to wait
	go decrementWaitGroupAfterCall(mailSent, "Pope", &wg)
	go decrementWaitGroupAfterCall(mailReceived, "Pope", &wg)
	wg.Wait() // will halt the further execution until WaitGroup counter reaches 0

	/*
		when we run goroutines they can't communicate with other goroutines
		// hence we use channels which act as a communication link between two or more go routines
		// channels  have type in golang like int,float,string
		// channels can send data and also can receive the data
		// once the channel is closed you cannot send or receive data from the channel
	*/

	// creating a channel (chan) using make() function
	chantChannel := make(chan string)
	// running chant as a goroutine
	go chant("Om", uint(5), chantChannel)

	// an anonymous function which executes as soon as it is declared
	func() {

		// to receive all the data from the channel we encapsulate the statement in the loop
		var chantReceived string
		var open bool
		// channel returns us two parameters one is
		// data and other is boolean which tells if the channel is open or not
		for {
			chantReceived, open = <-chantChannel
			if !open {
				break
			}
			fmt.Println(chantReceived)
		}
	}()

	// the call to the below function chant() function
	// will create a go routine with name chant and will send data to the channel of string
	// but the channel is closed so it will cause panic or generate no output
	// go chant("Mani", uint(5), chantChannel) // hence commented
	printChant(chantChannel) // no chants will be printed as the receiver is listening on closed channel

	/*
		Combination of waitGroup and channel
		case arise to use same channel in different goroutine
		 or what if we have multiple senders and only one receiver
	*/

	var wgc sync.WaitGroup
	wgc.Add(4)
	chantChannel3 := make(chan string)
	go decrementWaitGroupAfterChant(sendChant, "om", 5, chantChannel3, &wgc)
	go decrementWaitGroupAfterChant(sendChant, "mani", 5, chantChannel3, &wgc)
	go decrementWaitGroupAfterChant(sendChant, "padme", 5, chantChannel3, &wgc)
	go decrementWaitGroupAfterChant(sendChant, "hum", 5, chantChannel3, &wgc)
	go closeChannel(chantChannel3, &wgc)
	for chant := range chantChannel3 {
		fmt.Printf(" %v \t", chant)
	}
	// what will happen if we don't close the channel ?

	//  buffered channels
	bufferedChan := make(chan int, 4)
	go writeToBuffer(bufferedChan)
	time.Sleep(3 * time.Second)
	for v := range bufferedChan {
		fmt.Println("read value from ", v, " from ch")
		time.Sleep(2 * time.Second)
	}

	var fastChannel chan string = make(chan string)
	slowChannel := make(chan string)
	go replyFast(fastChannel)
	go replySlow(slowChannel)
	consumeBySwitch(fastChannel, slowChannel)
	fmt.Println("CLOSED !!!")

}

// end of main

func decrementWaitGroupAfterCall(mailType func(string), recipientName string, wg *sync.WaitGroup) {
	mailType(recipientName)
	wg.Done() // decrements the waitGroup counter by 1
}

func chant(mantra string, noOfTimes uint, chantChannel chan string) {
	var i uint
	for i < noOfTimes {
		chantChannel <- mantra
		i++
	}
	close(chantChannel)
}

func printChant(chantChannel chan string) {
	// we can also use range function on channel
	// by using range function we can skip if the channel is open or not
	for msg := range chantChannel {
		fmt.Println(msg)
	}
}

func sendChant(mantra string, noOfTimes uint, chantChannel chan string) {
	var i uint
	for i < noOfTimes {
		chantChannel <- mantra
		i++
	}
}

func decrementWaitGroupAfterChant(sendChant func(string, uint, chan string), chant string,
	noOfTimes uint, channel chan string, wgn *sync.WaitGroup) {
	defer wgn.Done()
	sendChant(chant, noOfTimes, channel)
}

func closeChannel(channel chan string, wgn *sync.WaitGroup) {
	wgn.Wait()
	close(channel)
}

// BUFFERED CHANNELS IN GO

// by default channels are un-buffered that
// which means a channel will only send data if it has a receiver at the end to consume the data

// so that why buffered channel were created which can store some values without a corresponding receiver at the other end to consume the data
// buffered channels are only blocked buffered capacity is full
func writeToBuffer(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("successfully wrote ", i, " to ch")
	}
	close(ch)
}

// using select case to consume channel which is ready first

func replyFast(channel chan string) {
	for i := 1; i <= 2; i++ {
		time.Sleep(time.Millisecond * 200)
		channel <- "in 200 ms"
	}
	close(channel)
}

func replySlow(channel chan string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Millisecond * 500)
		channel <- "in 2 s"
	}
	close(channel)

}

func consumeBySwitch(chan1, chan2 chan string) {
	for {
		select {
		case msg1, chan1Open := <-chan1:
			fmt.Println("successfully replied ", msg1)
			if !chan1Open {
				chan1 = nil
			}
		case msg2, chan2Open := <-chan2:
			fmt.Println("successfully replied ", msg2)
			if !chan2Open {
				chan2 = nil
			}
		}
	}
	/*
		The select statement is designed to wait on multiple channel operations
		The select statement has the following feature
		Random Selection: If multiple cases are ready to proceed, one of them is chosen at random.
		This ensures fairness and prevents goroutine starvation.

		Non-blocking with Default Case: The select statement can be made non-blocking by including a default case.
		If no other cases are ready, the default case is executed, preventing the select statement from blocking indefinitely

		but when the channel is closed the zero value for the channel's type and a false boolean value.
		When a channel is closed, attempting to receive from it will not block indefinitely.
		instead, it will immediately return the zero value for the channel's type, along with a boolean indicating that the channel is closed.

	*/
}
