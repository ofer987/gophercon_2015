# Solving the Go Latency Problem

By Rick Hudson, Google

- Go GC
  - Will be released in 1.5
- Go making a bet
  - More cores = better software
- Problem
  - GC pauses too often
- Less throughput for reduced GC latency
- Memory allocation strategies inspired by malloc from C
  - Hope to attract C developers
