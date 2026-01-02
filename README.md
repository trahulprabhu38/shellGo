# shellGo

A POSIX-compliant shell implementation written in Go, built as part of the [CodeCrafters](https://codecrafters.io) "Build Your Own Shell" challenge.

## Overview

shellGo is a basic command-line shell that demonstrates fundamental shell concepts including command parsing, built-in command execution, and a REPL (Read-Eval-Print Loop) interface.

## Features

- **Interactive REPL**: Continuous command-line interface with `$` prompt
- **Built-in Commands**:
  - `exit` - Exit the shell
  - `echo <args>` - Print arguments to stdout
  - `type <command>` - Display information about command type
- **Error Handling**: Graceful handling of unknown commands and input errors

## Prerequisites

- Go 1.25.0 or higher

## Installation

Clone the repository and build the project:

```bash
git clone <repository-url>
cd shellGo
go build -o shellgo ./app
```

## Usage

Run the shell:

```bash
./shellgo
```

Or use the provided shell script:

```bash
./your_program.sh
```

### Example Session

```bash
$ echo Hello, World!
Hello, World!

$ type echo
echo is a shell builtin

$ type ls
ls: not found

$ exit
```

## Project Structure

```
shellGo/
├── app/
│   └── main.go          # Main shell implementation
├── codecrafters.yml     # CodeCrafters configuration
├── go.mod              # Go module definition
├── your_program.sh     # Shell execution script
└── README.md           # This file
```

## Implementation Details

### Command Processing

The shell operates in a continuous loop:
1. Display prompt (`$`)
2. Read user input
3. Parse and execute command
4. Repeat

### Built-in Commands

Built-in commands are handled directly by the shell rather than being executed as external programs:

- **exit**: Terminates the shell process
- **echo**: Prints arguments with whitespace trimming
- **type**: Identifies whether a command is a shell built-in or external program

## Development

### Running Locally

```bash
go run ./app/main.go
```

### Testing

This project is part of the CodeCrafters challenge. To run tests:

```bash
# Follow CodeCrafters testing instructions
```

## Roadmap

Potential future enhancements:

- [ ] PATH resolution for external commands
- [ ] Command execution (fork/exec)
- [ ] Piping and redirection
- [ ] Environment variable support
- [ ] Command history
- [ ] Tab completion
- [ ] Signal handling (Ctrl+C, Ctrl+Z)
- [ ] Job control

## Contributing

This is a learning project, but suggestions and improvements are welcome. Feel free to:

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Open a pull request

## License

This project is part of CodeCrafters educational content.

## Acknowledgments

Built as part of the [CodeCrafters](https://codecrafters.io) "Build Your Own Shell" challenge.
