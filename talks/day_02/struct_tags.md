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
- [MustHaveTag](https://github.com/evergreen-ci/tags)
- [A More Unusual Mapping Case](https://github.com/mitchellh/mapstructure)
- [Command-line Configuration](https://github.com/jessevdk/go-flags)
- [Default Values](https://github.com/mcuadros/go-defaults)
  - Simplify initializers
  - Initializations logic live inside type definitions
- [Validation](https://github.com/asaskevich/govalidator)
  - Specify validators to be applied to a field
  - Validation specified declaratively
- Do not do
  - `interface{}` with type bounds

#### Namespacing

- Use `go vet` to catch errors
- [http://golang.org/pkg/reflect/#StructTag.Get](http://golang.org/pkg/reflect/#StructTag.Get)
