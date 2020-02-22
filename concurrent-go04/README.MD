# Channels

Channels support coordinate and share memory
goroutine => mutex/channel => goroutine

# Creating Channels
create a channel
ch := make(chan int)
// create a buffered channel
ch := make(chan int, 5)

# Channel Types
ch := make(chan int) // created channels are always bidirectional
func myFunction(ch chan int) {...}  // bidirectional channel
func myFunction(ch chan<- int) {...}  // send-only channel
func myFunction(ch <-chan int) {...}  // receive-only channel

# Control Flow
If statements
For loops
Select statements
ch1 := make(chan int)
ch2 := make(chan string)
select {
    case i := <-ch1:
        ...
    case ch2 <- "Hello":
        ...
    default:
        // use default case for non-blocking select
}