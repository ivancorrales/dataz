# dataz

Dataz is a library that embraces the Generics to simplify the way to work with data structures.

## Pre-requiites

Install Go v1.18 or higher

## Overview

Dataz provides a rich and growing set of utilities to work with data structures.

- Functional approach to deal with arrays.

## Installing

Using Dataz is straighforward. First, use go get to install the latest version of the library.

```bash
go get -u github.com/ivancorrales/dataz@latest
```

Next, include Dataz in your application:

```bash
import "github.com/ivancorrales/dataz"
```

## Usage

See below the list of provided utilities.

### Arrays

Provided operations with arrays:

- **Chunk:** It creates an array of elements split into groups the length of size. If array can't be split evenly, the final chunk will be the remaining elements.

```go
input:=[]int{2,3,-1,3,-2,6}
size:=4
output:= Chunk(input, size)

// output -> {{2,3,-1,3},{-2,6}}
```

- **Concat:** It creates a new array concatenating array with any additional arrays and/or values.

```go
input:[][]string{{"a", "b"}, {}, {"c"}, {"d"}},
output:= Concat(input)

// output -> {"a","b","c","d"}
```

- **Drop:** It removes the elements in the provided indexes.

```go
input:=   []string{"home", "port", "table", "window"}
indexes:= []int{1, 2}

output:=  Drop(input,indexes)

// output -> []string{"home", "window"}
```

- **DropWhile:** It drops the elements in the array until one of the elemnts doesn't match with the given criteria.

```go
input:=   []int{2,3,-1,3,-2,6}

output:=  DropWhile(input,func(v int)bool{return v>0})

// output -> []int{2,3}
```

- **Duplicates:** It returns a list with duplicated elements in the input.

```go
input:=   []int{2,3,-1,3,-2,6}

output:=  Duplicates(input)

// output -> []int{3}
```

- **Filter:** It removes those elements that doesn't match with the given criteria.

```go
input:=   []int{2,3,-1,3,-2,6}

output:=  Filter(input, func(v int)bool{return v>2})

// output -> []int{3,3,6}
```

- **ForEach:** It performs an action for each element in the input array.

```go
acc:=0
input:=[]int{2,3,-1,3,-2,6}
ForEach(input, func(v int){acc+=v})

// acc -> 11
```

- **Reverse:** It return the reversed inpur array

```go
input:=[]int{2,3,-1,3,-2,6}
output:= Reverse(input)

// output -> {6,-2,3,-1,3,2}
```






## License

Dataz is released under the Apache 2.0 license. See [LICENSE](LICENSE)
