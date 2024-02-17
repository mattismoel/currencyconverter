# What is this?

This project is a simple currency converter CLI tool written in Golang. It makes use of the free currency exchange API [Frankfurter](https://frankfurter.app).


# Usage

The command below builds the Go program and creates a binary executable file `currencyconverter` in the working directory.

```bash
$ go build -o currencyconverter ./cmd/*
```

The code can also be directly run with the following command:
```bash
$ go run ./cmd/main.go <amount> <from_currency> <to_currency>
```
Or the newly built binary:
```bash
$ ./currencyconverter <amount> <from_currency> <to_currency>
```
