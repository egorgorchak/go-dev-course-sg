package main

func main() {
	c := make(chan string)
	// c := make(chan string, 1)
	c <- "Hi there!"
}

//This is correct, you cannot send anything to a channel if there's no other go routine to receive it, unless you make a buffered channel like so: c := make(chan string, 1), where 1 is the number of items in the buffer.
