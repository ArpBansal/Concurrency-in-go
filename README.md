# Concurrency-in-Go

This repository contains three Go files that demonstrate the differences between sequential execution, unfair threading, and fair threading.

## Files

1. **sequential.go**: Showcases sequential execution.

time of execution for my system - 1m32.108311035s
2. **unfair_threads.go**: Demonstrates unfair threading.

time of execution for my system - 17.330201271s
3. **fair_threads.go**: Illustrates fair threading.

time of execution for my system - 14.074594448s
## Usage

To run any of the examples, use the `go run` command followed by the file name. For example:

```sh
go run {folder}/{filename}
```

## Description

- **Sequential Execution**: This example runs tasks one after another without any concurrency.
- **Unfair Threads**: This example demonstrates a scenario where threads are not given equal opportunity to execute.
- **Fair Threads**: This example ensures that all threads get a fair chance to execute.

Feel free to explore the code and understand the differences in execution patterns.

## Requirements

- I used go 1.23

## License

This project is licensed under the MIT License.