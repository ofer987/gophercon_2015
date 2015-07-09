# [The Many Faces of Struct Tags](http://www.gophercon.com/talks/struct-tags)

By
- [Kyle Erf](https://github.com/3rf), Software Engineer at [Flatiron Health](http://www.flatiron.com/)
- [Sam Helman](https://github.com/shelman), Kernel Tools Engineer at [MongoDB](https://www.mongodb.org/)

#### Struct Tag

- An arbitrary unicode string value attach to a struct field
- Tag fields of a struct
- Use cases
  - Serialize to and fro JSON
    - `encoding/json`
  - XML
  - BSON
  - SQL ORMS

#### Namespacing

- Use `go vet` to catch errors
- [http://golang.org/pkg/reflect/#StructTag.Get](http://golang.org/pkg/reflect/#StructTag.Get)
