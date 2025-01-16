# Go Deadlock Example

This repository demonstrates a common deadlock scenario in Go involving mutexes and channels.  Ten goroutines acquire a mutex, send a value to a channel, and then release the mutex.  A separate goroutine receives from the channel.

The program may deadlock if the receiver consumes all available data from the channel before all the sending goroutines have completed.