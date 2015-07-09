# [The Evolution of Go](http://www.gophercon.com/talks/evolution-of-go)

By [Robert Griesemer](https://en.wikipedia.org/wiki/Robert_Griesemer), Engineer at [Google](http://www.google.com)

#### Inpiration

- Lego
- "Hints on Programming Language Design", by C.A.R. Hoare
- "Everything You Wanted to Know About Programming langauges But Were Afraid to Ask", by C.A.R. Hoare
- 'The task of the programming language designer is **consolidation not innovation**' (Hoare, 1973)
- Simplicity
- One way to write a construct
- Expressing algorithms
- Not the type system

#### Predecessors

- Algol 60
  - Pascal (1968 - 1970)
  - C (1969 - 1973)
- The heritage of GO
  - Oberon
  - C
- Go joins Pascal and C back together

#### Interfaces

- Interfaces allow *messaging* between objects like in Smalltalk and Ruby

#### Generics

- Also known as Parametric polymorphism
- Non-orthogonal feature
  - Interacts with other features
  - Increases complexity
- Missing in Go
- Rarely missed by Go programmers
- Will not be implemented anytime soon
  - Designers do not understand it yet (how to implement it)

#### Specification

- Initially designed from 2007 to 2009
- Go/types spec arriving in version 1.5
- Tools for ensuring language evolves properly
  - Does not leave obsolete code
  - `gofmt`
    - language changes
  - `gofix`
    - API changes

#### Success

- What Characteristics does a language require to be successful?

- Clear target
- Solid implementation
  - Language
  - Libraries
  - Tools
- Marketing readiness
- Technological breakthrough
- Leanguage features without competitors
- Marketing (rarely)
  - Like Java
  - However, Sun is no longer alive
- Programming languages need ten years to become established
- `go build` allows for compilation without a make file
