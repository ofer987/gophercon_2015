# [Preventing Deadlocks and Leaks in Go](http://www.gophercon.com/talks/deadlocks-leaks)

By [Richard Fliam](https://github.com/rfliam), Software Engineer at [Comcast](http://www.xfinity.com/)

- Use CSP (Communicating Sequential Processors)
  - By C.A.R. Hoare, the designer of CSP
- Not the same as designing for object-oriented design
- Focus on the flow of data
  - Do not focus on the flow of control

#### Draw a diagram

- Look for problems
  - God Process
    - Too busy
    - Too many types
    - Too many processes
  - Rat's Nest
    - High connectivity
    - Muddled communication
  - Cycles
    - Possible deadlocks

- Create pipelines
  - Unidirectional flows of data
    - Cannot deadlock
  - Pub/Sub is not unidirectional
    - So can deadlock
    - Turn Pub/Sub to unidirectional:
      - context
      - unsubscribe
      - publisher
      - event
      - subscriber
- Do not let exiting leak
  - Do not close goroutine irresponsibly
  - Close with Context.Done()
  - Do not block

- Collecting Results
  - Cancel
  - Collect
  - Exit

#### Tips

- Data Flow
- Drawing
- Piplines
- Existing

- Stop tickers
- Close contexts
  - `defer cancel()`
- Defer properly
  - `defer f.Close()`
- Finish your goroutines
  - So the runtime can collect it
