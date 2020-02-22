# Channels

Channels support coordinate and share memory
goroutine => mutex/channel => goroutine

# Creating Channels
create a channel
ch := make(chan int)
// create a buffered channel
ch := make(chan int, 5)