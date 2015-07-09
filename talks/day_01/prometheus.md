# [Prometheus: Designing and Implementing a Modern Monitoring Solution in Go](http://prometheus.io)

By [Björn Rabenstein](https://github.com/beorn7), Production Engineer at [SoundCloud](https://soundcloud.com/)

- package `sync/atomic`
  - 25x faster than mutex counter
- Use -benchmem
 - To detect allocation churn
 - go test -bench=. -cpu=1,4,16 -benchmem
- Escape analysis:
   - go test -gcflags=-m -bench=Something
- Use pprof
  - import "net/http/pprof"
- Use cgo

[Performance Tales](http:/jmoiron.net/blog/go-performance-tales)
