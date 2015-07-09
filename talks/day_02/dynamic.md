# [Go Dynamic Tools](http://www.gophercon.com/talks/dmitry-vyukov)

By [Dmitry Vyukov](https://github.com/dvyukov), Engineer at [Google](http://www.google.com)

- [go-fuzz](https://github.com/dvyukov/go-fuzz)
  - Available in Go 1.5
  - go get github.com/dvyukov/go-fuzz

  - go-fuzz-build github.com/dvyukov/go-fuzz/examples/go
  - go-fuzz -bin=go-fuzz.zip -workdir=some_dir

- Find race conditions
  - Use `-race` flag
    - Go run
    - Go test
    - etc.

- Execution Tracer
  - Use with `-trace`

- [versifier](https://libraries.io/go/github.com%2Fdvyukov%2Fgo-fuzz%2Fgo-fuzz%2Fversifier)
