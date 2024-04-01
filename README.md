# GoSh

[![LICENSE](https://img.shields.io/github/license/mattheworford/gosh.svg?style=flat-square)](https://github.com/mattheworford/gosh/LICENSE)

GoSh is a basic shell implementation in Go, providing a simple command-line interface for interacting with your operating system.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mattheworford/gosh.git
   ```

2. Navigate to the project directory:

   ```bash
   cd gosh
   ```

3. Build the executable:

   ```bash
   go build
   ```

### Usage

Run the executable:

```bash
./gosh
```

You should now be in the interactive shell prompt.

## Features

- **Basic Commands:** Enter system commands like `ls`, `cat <filename>`, etc.
- **Built-in Commands:**
  - `cd [directory]`: Change the current working directory.
  - `pwd`: Print the working directory.
  - `history`: List the command history.
  - `exit`: Exit the shell.
- **Command Piping:** Chain commands together with the `|` operator to utilize the output of one as input for another.
- **History Persistence**: Automatically store history to and load history from a `.gosh_history` file in your machine's home directory.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This project is inspired and guided by ["Write Your Own Shell"](https://codingchallenges.fyi/challenges/challenge-shell) from [John Crickett](https://uk.linkedin.com/in/johncrickett)'s
[Coding Challenges](https://codingchallenges.fyi/).
