# Sets Calculator

Program should print the result on standard output: sorted integers, one integer in a line.

Grammar of calculator is given:

```
expression := "[" operator N sets "]" 
sets := set | set sets
set := file | expression
operator := "EQ" | "LE" | "GR"
```

- "file" is a file with sorted integers, one integer in a line.

- "N" is a positive integer

Meaning of operators:

- `EQ` - returns a set of integers which consists only from values which exists in exactly N sets - arguments of operator
- `LE` - returns a set of integers which consists only from values which exists in less then N sets - arguments of operator
- `GR` - returns a set of integers which consists only from values which exists in more then N sets - arguments of operator

```shell script
# make example_1
$ ./scalc [ LE 2 a.txt [ GR 1 b.txt c.txt ] ] 
1
4

# make example_2
$ ./scalc [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ] 
2
3
```

## How to run

> At first, you need an installed [Go 1.14+](https://golang.org/doc/install)

```shell script
go build .

./scalc [ LE 2 a.txt [ GR 1 b.txt c.txt ] ] 
```

Or 

```shell script
make build

make example_1

make example_2
```
