# Concurrency

- Managing a lot of things at once
- Goroutine is a function that is scheduled to run independently of other goroutines
- Single logical processor
  - Assigned a single processing thread
  - Asigned a local run queue
  - Blocking goroutines are placed off the queue
  - By default 10,000 goroutines on one thread
- Best mentality during development:
  - Goroutines run in parallel, each on its own processor
  - Regardless of how many physical processors exist on machine
  - Goroutines are scheduled in and out by the scheduler every ten milliseconds
- Convention:
  - Start cleanly
  - Shutdown cleanly
- Tooling support:
  - -race flag to detect race conditions
  - Used with `go build` and `go test`
- `go` keyword in front of function
  - Schedule the function to be managed by the scheduler
- Program panics, i.e., `fatal error`
  - if all goroutines are deadlocked

### Mutexes and Semaphores

- sync.WaitGroup

### Channels

- Provide
  - Guarantee
  - Responsibility
- `chan` is a reference type
- State
  - Closed
    - Every receive returns
    - Cannot send
    - Can receive
  - Non-closed
- defer execute after the return statement
  - Guaranteed to be called, even if the function panics
  - Can stop a panic from bubbling up the stack with `recover`
  - Do not hide panic attacks, so log them
    - use the logger from standard library
  - Should not expect to receive panics so do not code against them

#### Unbuffered Channels

- Send and receive

- Advantages
  - Synchronization thanks to blocking
  - Guarantee

#### Buffered Channels

- No blocking unless capacity is full

