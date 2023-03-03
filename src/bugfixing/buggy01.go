/*
Buggy:
package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	ch <- "Hello world!"        <---- PROBLEM: since ch is unbuffered, the sending blocks thread until something recieves.
	fmt.Println(<-ch)						   but it wont be recieved since the reciever is on same thread.
}
*/

/*
Simple solution is to make `ch` buffered. Eg. "ch := make(chan string, 1)". But that is lame
Chad solution is to make the sending in a seperate goroutine
*/

package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	go func(c chan string) { c <- "Hello world!" }(ch)
	fmt.Println(<-ch)
}

/*
Now the sending still blocks, but it blocks a seperate thread.
*/
