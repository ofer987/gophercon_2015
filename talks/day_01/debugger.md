# [Debugger](https://github.com/derekparker/delve)

By [Derek Parker](http://derkthedaring.com/), Engineer at [Hashrocket](http://hashrocket.com/)

## [Delve](https://github.com/derekparker/delve) 

- `dlv run`
  - compiles code with optimizations disabled
  - starts program
  - attaches to process
- `dlv test`
  - compiles a test binary
  - runs tests
- `dlv attach <pid>`
  - warning: program might be compiled with optimizations
- Support for editor integration
  - JSON-RPC
- Readline support!

#### Commands

- `break`
- `trace`
- `continue`
- `step`
- `next`
- `thread <tid>`
- `restart`
- `threads`
- `goroutines`
- `breakpoints`
- `info func /regex/`

- Insert breakpoint using `runtime.Breakpoint` within source code
