<p align="center">
  <img src="logo.svg" width="40%" border="0" alt="avo" />
  <br />
  <a href="https://app.shippable.com/github/mmcloughlin/avo/dashboard"><img src="https://api.shippable.com/projects/5bf9e8f059e32e0700ec360f/badge?branch=master" alt="Build Status" /></a>
  <a href="http://godoc.org/github.com/mmcloughlin/avo"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg" alt="GoDoc" /></a>
</p>

<p align="center">High-level Golang x86 Assembly Generator</p>

`avo` aims to make high-performance Go assembly easier to write, review and maintain. It's a Go package that presents a familiar assembly-like interface, together with features to simplify development without sacrificing performance:

* `avo` programs _are_ Go programs: use **control structures** for assembly generation
* **Register allocation**: write your kernels with **virtual registers** and `avo` assigns physical registers for you
* Automatic **parameter load/stores**: ensure memory offsets are always correct even for complex data structures
* Generation of **stub files** to interface with your Go package

Inspired by the [PeachPy](https://github.com/Maratyszcza/PeachPy) and [asmjit](https://github.com/asmjit/asmjit) projects.

_Note: `avo` is still in an experimental phase. APIs subject to change._

## Install

Install `avo` with `go get`:

```
$ go get -u github.com/mmcloughlin/avo
```

## Quick Start

`avo` assembly generators are pure Go programs. Let's get started with a function that adds two `uint64` values.

[embedmd]:# (examples/add/asm.go)
```go
// +build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
)

func main() {
	TEXT("Add", "func(x, y uint64) uint64")
	Doc("Add adds x and y.")
	x := Load(Param("x"), GP64v())
	y := Load(Param("y"), GP64v())
	ADDQ(x, y)
	Store(y, ReturnIndex(0))
	RET()
	Generate()
}
```

You can `go run` this code to see the assembly output. To integrate this into the rest of your Go package we recommend a [`go:generate`](https://blog.golang.org/generate) line to produce the assembly and the corresponding Go stub file.

[embedmd]:# (examples/add/add_test.go go /.*go:generate.*/)
```go
//go:generate go run asm.go -out add.s -stubs stub.go
```

After running `go generate` the [`add.s`](examples/add/add.s) file will contain the Go assembly.

[embedmd]:# (examples/add/add.s)
```s
// Code generated by command: go run asm.go -out add.s -stubs stub.go. DO NOT EDIT.

// func Add(x uint64, y uint64) uint64
TEXT ·Add(SB), $0-24
	MOVQ	x(FP), AX
	MOVQ	y+8(FP), CX
	ADDQ	AX, CX
	MOVQ	CX, ret+16(FP)
	RET
```

The same call will produce the stub file [`stub.go`](examples/add/stub.go) which will enable the function to be called from your Go code.

[embedmd]:# (examples/add/stub.go)
```go
// Code generated by command: go run asm.go -out add.s -stubs stub.go. DO NOT EDIT.

package add

// Add adds x and y.
func Add(x uint64, y uint64) uint64
```

See the [`examples/add`](examples/add) directory for the complete working example.

## Examples

### Slice Sum

Sum a slice of `uint64`s:

[embedmd]:# (examples/sum/asm.go /func main/ /^}/)
```go
func main() {
	TEXT("Sum", "func(xs []uint64) uint64")
	Doc("Sum returns the sum of the elements in xs.")
	ptr := Load(Param("xs").Base(), GP64v())
	n := Load(Param("xs").Len(), GP64v())
	s := GP64v()
	XORQ(s, s)
	LABEL("loop")
	CMPQ(n, operand.Imm(0))
	JE(operand.LabelRef("done"))
	ADDQ(operand.Mem{Base: ptr}, s)
	ADDQ(operand.Imm(8), ptr)
	DECQ(n)
	JMP(operand.LabelRef("loop"))
	LABEL("done")
	Store(s, ReturnIndex(0))
	RET()
	Generate()
}
```

### Parameter Load/Store

`avo` provides deconstruction of complex data datatypes into components. For example, load the length of a string argument with:

[embedmd]:# (examples/args/asm.go go /.*TEXT.*StringLen/ /Load.*/)
```go
	TEXT("StringLen", "func(s string) int")
	strlen := Load(Param("s").Len(), GP64v())
```

Index an array:

[embedmd]:# (examples/args/asm.go go /.*TEXT.*ArrayThree/ /Load.*/)
```go
	TEXT("ArrayThree", "func(a [7]uint64) uint64")
	a3 := Load(Param("a").Index(3), GP64v())
```

Access a struct field (provided you have loaded your package with the `Package` function):

[embedmd]:# (examples/args/asm.go go /.*TEXT.*FieldFloat64/ /Load.*/)
```go
	TEXT("FieldFloat64", "func(s Struct) float64")
	f64 := Load(Param("s").Field("Float64"), Xv())
```

Component accesses can be arbitrarily nested:

[embedmd]:# (examples/args/asm.go go /.*TEXT.*FieldArrayTwoBTwo/ /Load.*/)
```go
	TEXT("FieldArrayTwoBTwo", "func(s Struct) byte")
	b2 := Load(Param("s").Field("Array").Index(2).Field("B").Index(2), GP8v())
```

Very similar techniques apply to writing return values. See [`examples/args`](examples/args) and [`examples/returns`](examples/returns) for the full suite of examples.

### Real Examples

* **[fnv1a](fnv1a):** [FNV-1a](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV-1a_hash) hash function.
* **[dot](dot):** Vector dot product.
* **[geohash](geohash):** Integer [geohash](https://en.wikipedia.org/wiki/Geohash) encoding.
* **[sha1](sha1):** [SHA-1](https://en.wikipedia.org/wiki/SHA-1) cryptographic hash.
* **[stadtx](stadtx):** [`StadtX` hash](https://github.com/demerphq/BeagleHash) port from [dgryski/go-stadtx](https://github.com/dgryski/go-stadtx).

## Contributing

Contributions to `avo` are welcome:

* Feedback from using `avo` in a real project is incredibly valuable.
* [Submit bug reports](https://github.com/mmcloughlin/avo/issues/new) to the issues page.
* Pull requests accepted. Take a look at outstanding [issues](https://github.com/mmcloughlin/avo/issues) for ideas (especially the ["good first issue"](https://github.com/mmcloughlin/avo/labels/good%20first%20issue) label).

## License

`avo` is available under the [BSD 3-Clause License](LICENSE).
