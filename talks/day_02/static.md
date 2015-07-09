# [Static Code Analysis Using SSA](http://www.gophercon.com/talks/static-code-analysis)

By [Ben Johnson](https://www.github.com/benbjohnson), Go Developer

- Implementer of [go-raft](https://github.com/goraft/raft)

- [golang.org/x/tools/go/ssa](http://golang.org/x/tools/go/ssa)
  - Written by Alan Donovan
    - Who also co-wrote the book **Go Programming Language**

#### Tools written based on go/ssa

- [SafeSQL](https://github.com/stripe/safesql)
  - Prevents SQL injection
- [Oracle](https://godoc.org/golang.org/x/tools/cmd/oracle)
  - Given a variable, what it stored, and how
    - Works with interfaces
  - And more
- [gorename](https://godoc.org/golang.org/x/tools/cmd/gorename)
  - Like search-and-replace but without errors
  - To rename identifiers
    - Declarations
    - References
  - Prevents
    - Ambiguitity
    - Shadowing conflicts
      - Variable already in use
