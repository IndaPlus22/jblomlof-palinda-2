/*package main

import (
	"fmt"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	ch := make(chan int)
	go Print(ch)					<---- Problem(s): we start a new goroutine here, which recieves from main
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)						<---- We exit main without thinking about if our routine has finished or not.
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}
*/

/*
cleanest solution would be to refractor the program - so that main recieves and not sends.
but i cba to refractor it all so a quick fix is to make an additional channel in whichs main recieves to.
*/

package main

import (
	"fmt"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	ch := make(chan int)
	doneChecker := make(chan int)
	go Print(ch, doneChecker)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	//blocks until other goroutine is finished
	<-doneChecker
}

// Print prints all numbers sent on the recieveing channel.
// The function returns when the channel is closed and sends 0 to the second channel
func Print(ch <-chan int, doneChecker chan<- int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	//send back to main that we are
	doneChecker <- 0
}

/*
This now works since main blocks to recieve from print, thus letting print finish.
*/
