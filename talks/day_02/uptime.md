# [Uptime: Building Resilient Services With Go](http://www.gophercon.com/talks/uptime/)

By [Blake Caldwell](https://github.com/wblakecaldwell), Developer at [Fog Creek Software](https://www.fogcreek.com/) | Kentik

- Cut something by half!
  - Wow

#### Channel Axioms

1. A send to a nil channel blocks forever
2. A receiver from a nil channel blocks forever
3. A send to a closed channel pnaics
4. A receive from a closed channel returns the zero value immediately

#### Recover From Panics

- Only recover if you are sure it is okay
- Panic recovery is for current goroutine
- Log the stack trace

#### Go Race Detector

- Use in unit tests
- Environments
  - Development
  - Test
- Use the `-race` flag
  1. Test
  2. Run
  3. Build
  4. Install
