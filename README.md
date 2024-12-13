[![Go Reference](https://pkg.go.dev/badge/chipaca.com/intmath.svg)][package documentation]

# intmath
[intmath] implements some useful integer math functions. Think of it as the
package that complements [math/bits] to have a wholly usable [math] but for
integers.

You can read the [package documentation]. You can also read a [rambly blogpost]
about it. Your choice!

[math]: https://pkg.go.dev/math
[math/bits]: https://pkg.go.dev/math/bits
[intmath]: https://github.com/chipaca/intmath
[package documentation]: https://pkg.go.dev/chipaca.com/intmath
[rambly blogpost]: https://chipaca.com/en/2024/01/integer-math/

---

In the tests, you'll find

* "basic" tests that sanity check form a handful of values, comparing
  the result to known values.
* "quick" tests (using testing/quick) that compare the result with
  slower, assumed-correct implementations.  Some of these are _not_
  quick, despite the name; if you run the tests without `-short`
  you'll probably need to add a bigger `-timeout`.
* benchmarks that compare the functions to the slower versions.
