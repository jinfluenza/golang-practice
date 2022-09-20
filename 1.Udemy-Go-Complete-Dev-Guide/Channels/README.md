# Go routine/channel

- Makes it un in multithreading fashion
- Allows asynchronous call
- To enable multi threading use `go FUNCTION()`, but this doesn't manage the subprocesses ran by the main process
- Go scheduler runs one routine util it finishes or makes a blocking call
- Concurrency is not parallelism 
- Concurrency: we have multiple threads executing code. If one threads blcoks another one is picked up and worked on
  - Program has ability to run multiple routine
- Parallelism - multiple threads executed at the exact same time. Requires multiple CPU cores
  - while one core runs and go routine, other core can run other go routine
- Channel is like a middleware that can control the threads
- Routine is waiting for it to hear something back (put to sleep), blocking channel
- Lambda can be used like this:
```go

func() {

} () //to invoke you put () at the end

```
- pass by value -> reference.
- passing the value in argument creates a new value in the memory so that the main routine can replace the value.