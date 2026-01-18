# shellGo

A simple shell implementation in Go, built as part of the [CodeCrafters](https://codecrafters.io) "Build Your Own Shell" challenge.

## Features

- Interactive REPL with `$` prompt
- Built-in commands: `exit`, `echo`, `type`
- Error handling for unknown commands

## Quick Start

```bash
# Run directly
go run ./app/main.go

# Or build and run
go build -o shellgo ./app
./shellgo
```

## Example

```bash
$ echo Hello, World!
Hello, World!

$ type echo
echo is a shell builtin

$ type ls
ls: not found

$ exit
```

## Requirements

- Go 1.25+

fuckkkkk i didnt work on this today

