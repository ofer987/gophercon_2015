# [Embrace the Interface](http://www.gophercon.com/talks/embrace-the-interface)

By [Tomás Senart](https://github.com/tsenart), Author of [Vegeta](https://github.com/tsenart/vegeta)

- Use interfaces to separate concerns
  - Reduce the scope of each interface
  - Use single-method interfaces

#### Decorator Pattern

- `type Decorator func(Client) Client`

- e.g.,

func Logging(l *log.Logger) Decorator {
  return func(c Client) Client {
    return ClientFunc(func(r *http.Request) (*http.Response, error) {
      l.Printf("Hello")
      return c.Do(r)
    })
  }
}
