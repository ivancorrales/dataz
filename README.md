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

- **GroupBy:** It groups the items in the provided array by groups of N elements. N is provided as parameter

```go
input:=[]int{2,3,-1,3,-2,6}
by:=4
output:= GroupBy(input, by)

// output -> {{2,3,-1,3},{-2,6}}
```
- **Reverse:** It return the reversed inpur array

```go
input:=[]int{2,3,-1,3,-2,6}
output:= Reverse(input)

// output -> {6,-2,3,-1,3,2}
```






## License

Dataz is released under the Apache 2.0 license. See [LICENSE](LICENSE)
