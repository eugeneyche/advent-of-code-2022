# Advent of Code 2022

These are my solutions to problems in the Advent of Code 2022.

All solutions are written in Golang, because I am using AoC to practice this new language.
Most of the code will probably not be perfect, but please bear with me. 😊

## Running Solutions

Most solutions will contain 3 files:

- `main.go` - the source code
- `sample.txt` - the AoC sample input
- `input.txt` - my input for this question

In some cases, sample inputs may be split up between parts. For example:

- `sample_p1.txt` - the AoC sample input for part 1
- `sample_p2.txt` - the AoC sample input for part 2

To run the solutions, first compile the source code with `go build main.go`.

Execute the binary by passing in the path of the input as the first argument

```bash
$ ./main sample.txt # Runs sample input
$ ./main input.txt  # Runs actual input
```

This will run the solution for part 2 of that day's problem. To run the solution
for part 1, first find the following lines in `main.go`.

```go
_ = p1
p2()
```

Update those lines to the following:

```go
p1()
_ = p2
```

Finally, repeat the steps to compile and run the solution.

Some solutions may need special arguments passed in to execute. These arguments
are parameters that differ between the sample and real input. In those cases, a
`README.md` file will be provided with instructions on how to run sample code
and actual input for the two parts.
