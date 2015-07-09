# [Solving the Go Latency Problem](http://www.gophercon.com/talks/garbage-collection/)

By [Richard Hudson](https://www.linkedin.com/pub/richard-hudson/85/391/333), [Google](https://www.google.com)

- Go GC
  - Will be released in 1.5
- Go making a bet
  - More cores = better software
- Problem
  - GC pauses too often
- Less throughput for reduced GC latency
- Memory allocation strategies inspired by malloc from C
  - Hope to attract C developers
