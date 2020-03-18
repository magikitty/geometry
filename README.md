# Geometry

## Overview

This is a library containing structs and methods that are useful in geometry. This project is written in Go.

## Types included in this library

### Interface

The interface `Shape` is implemented by the shapes (Circle and Rectangle) in this library. The method set for `Shape` is:

```go
- Area
- Perimeter
```

`shapes.go` also contains two functions that use `Shape` as a parameter:

- `TotalArea`: returns the combined area of any number of shapes passed to it
- `TotalPerimeter`: returns the combined perimeter of any number of shapes passed to it

### Structs and methods

The structs and their corresponding methods contained in this library are:

```go
- Circle struct
  - Area
  - Perimeter (i.e. circumference)
  - Diameter
- Point struct
- Rectangle struct
  - Area
  - Length
  - Perimeter
  - Width
```

## Rounding decimals

This library also contains the function `RoundNumber`, which returns the given number rounded to a specified number of decimal places. You can use `RoundNumber` to round between 0-10 decimal places.

## How to initialize an instance of a struct using this library

Call the `Make` functions found in each struct's file to create new instances of the struct.

E.g. the `MakeRectangle` function located in `rectangle.go` returns a new rectangle.

## Use the geometry library

To import the library into a Go project: run `go get github.com/magikitty/geometry` in the terminal. Then add the library path `github.com/magikitty/geometry` to the `import` list in your Go project.

To clone the project: run `git clone https://github.com/magikitty/geometry.git` in the directory you want.

## My profile

You can check out the rest of my projects on my [github profile](https://github.com/magikitty).
