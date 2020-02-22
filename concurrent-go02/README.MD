
# Init module
go mod init hadwin.com/hello

# Run 
go run .

# Thread
Have own execution stack
Fixed stack space (around 1MB)
Managed by OS

# Goroutine
Have own execution stack
Variable stack space (starts @2 KB)
Managed by Go runtime