# [What Could Go Wrong](http://www.gophercon.com/talks/what-could-go-wrong)

By [Kevin Cantwell](https://github.com/kevin-cantwell), Lead Architect at [Timehop](http://timehop.com/)

- Use runtime.GoSched() to unblock the scheduler
  - Might not be necessary in newer versions?
- Declare errors as interfaces (types?)
- Wrap leaky defers in anonymous functions: `defer action`
  - Or, 'for { ... action }' with the defer
- Can redeclare predeclared identifiers
