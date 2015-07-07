## Slices

- Core data structure
- Slice is a reference type
- Only use cache for speed
  - Use slices
    - Use arrays underneath for continguous memory allocation
  - More continguous memory allocation = More performance
- Nil can be a slice
  - var data []string
  - Characteristics:
    - 0 for the length,
    - 0 for the address
  - Do not use the nil slice unless for readability
  - More idiomatic to use nil slice than the empty slice
    - If use the JSON marshaler then better to use empty slice because it returns empty object
- Convention
  - Do not share reference type values in Go, instead share the whole slice by sharing its header
  - No random numbers
- When appending new items to slice then slice grows
- When slice grows:
  - Create new slice
  - Copy values from old slice to new slice
  - Discard old slice if it is no longer referenced
- For range `for i_counter, item_variable := range slice_variable { ... }`
  - Works with slice, maps, channels
  - slice item is copied by value to `item_variable`
- Copy range of slice to new slice from `start_index` to `end_index`
  - `slice2 := slice1[2:4]`
- Copy three-index slice
  - `slice2 := slice1[2:4:4]`
  - The third index specifies capacity
  - Useful for allocating a new array when the length of slice2 exceeds its capacity
    - Otherwise it would modify items of slice1 (i.e., the original slice)
    - Code smell if third index (i.e., capacity) not the same as second index (i.e., `end_index`)
      - Otherwise the target slice would modify the source slice when target slice grows
- Deep copy for copying source slice to target slice
  - Use `copy` function

### Strings

- UTF-8 encoded
- `strings` are really `slices`
- `rune` is an alias for `integer` can be 1 - 4 bytes long
- Funny quote: 'An array in Go is just a slice waiting to happen'
