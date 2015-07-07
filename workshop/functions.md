# Functions

- Methods are functions where the first parameter is the receiver
- Functions
  - Declared without receiver
- Methods
  - Declared with receiver
  - Receiver should be declared something unique
    - Not `self` or `this`
    - Usually 1 - 3 characters in length
    - Either value (copied by value) or pointer (points to value) receiver
- The closer the variable is declared the less descriptive the name is because it has context
- Functions declared with a value receiver can be used with
  - Value receiver
  - Pointer receiver
- Functions defined with a pointer receiver can be used with
  - Pointer receiver

### Types

- `type duration int64`
  - `duration` is a new type

### Functions As First-Class Citizions

- A function is a reference type
  - It has a header
  - Context needed to execute
- Pass a function around
- Allows currying
- When calling a method as a variable, you can pass the receiver as the first argument

### Interfaces

- Focus on the behaviour of the program
  - Mentality on API design
- Convention:
  - Names of interfaces with one method (no arguements) end in 'er'
- Represent a reference type
  - Header: Type
  - Pointer: Address of concrete type value that implements the interface
- Values implement the interface
  - As opposed to types implementing interfaces in other languages
- Implement the interface on a **pointer receiver**
- iTables allow polymorphic behaviour by mapping a function to the interface type
- Composition
  - Outer types are promoted to implement the interface of the their inner types
  - The implementation of the outer types override the implementation of the inner types
