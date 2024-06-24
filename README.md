# github.com/codescalersinternships/Coreutils-EYAD_HUSSEIN

github.com/codescalersinternships/Coreutils-EYAD_HUSSEIN is a collection of Unix-like command line utilities implemented in Go.

## Project Structure

```bash
.
├── bin
├── pkg
├── README.md
└── src
    └── github.com/codescalersinternships/Coreutils-EYAD_HUSSEIN
        ├── cmd
        │   ├── cat
        │   │   └── main.go
        │   ├── echo
        │   │   └── main.go
        │   ├── env
        │   │   └── main.go
        │   ├── false
        │   │   └── main.go
        │   ├── head
        │   │   └── main.go
        │   ├── tail
        │   │   └── main.go
        │   ├── tree
        │   │   └── main.go
        │   ├── true
        │   │   └── main.go
        │   ├── wc
        │   │   └── main.go
        │   └── yes
        │       └── main.go
        ├── commands
        │   ├── cat.go
        │   ├── echo.go
        │   ├── env.go
        │   ├── false.go
        │   ├── head.go
        │   ├── tail.go
        │   ├── tree.go
        │   ├── true.go
        │   ├── wc.go
        │   └── yes.go
        ├── go.mod
        ├── models
        │   ├── command.go
        │   └── commands.go
        └── utils
            ├── generalUtils.go
            └── validateInput.go
```

- **bin/**: Contains executable binaries.
- **pkg/**: Holds package files.
- **src/github.com/codescalersinternships/Coreutils-EYAD_HUSSEIN/**: Source directory for the project.

## Commands Implemented

- **cat**: Concatenate files and print on the standard output.

  - Allowed flags: `-n`

- **echo**: Display line of text in standard output.

  - Allowed flags: `-n`

- **env**: List all environment variables.

  - No additional flags allowed.

- **false**: Do nothing, unsuccessfully.

  - No additional flags allowed.

- **head**: Output the first part of files.

  - Allowed flags: `-n`

- **tail**: Output the last part of files.

  - Allowed flags: `-n`

- **tree**: List contents of directories in a tree-like format.

  - Allowed flags: `-L`

- **true**: Do nothing, successfully.

  - No additional flags allowed.

- **wc**: Print newline, word, and byte counts for each file.

  - Allowed flags: `-l`, `-w`, `-c`

- **yes**: Output a string repeatedly until killed.
  - No additional flags allowed.

## Getting Started

### Installation

Clone the repository and navigate to `src/github.com/codescalersinternships/Coreutils-EYAD_HUSSEIN`.

### Building Commands

You can build each command separately. Below are the instructions for building and using each command.

#### cat

1. **Build**:

   ```bash
   go build -o cat cmd/cat/main.go
   ```

2. **Usage**:
   ```bash
   ./cat [flags] <file>...
   ```

#### echo

1. **Build**:

   ```bash
   go build -o echo cmd/echo/main.go
   ```

2. **Usage**:
   ```bash
   ./echo [flags] <text>...
   ```

#### env

1. **Build**:

   ```bash
   go build -o env cmd/env/main.go
   ```

2. **Usage**:
   ```bash
   ./env
   ```

#### false

1. **Build**:

   ```bash
   go build -o false cmd/false/main.go
   ```

2. **Usage**:
   ```bash
   ./false
   ```

#### head

1. **Build**:

   ```bash
   go build -o head cmd/head/main.go
   ```

2. **Usage**:
   ```bash
   ./head [flags] <file>...
   ```

#### tail

1. **Build**:

   ```bash
   go build -o tail cmd/tail/main.go
   ```

2. **Usage**:
   ```bash
   ./tail [flags] <file>...
   ```

#### tree

1. **Build**:

   ```bash
   go build -o tree cmd/tree/main.go
   ```

2. **Usage**:
   ```bash
   ./tree [flags] <directory>
   ```

#### true

1. **Build**:

   ```bash
   go build -o true cmd/true/main.go
   ```

2. **Usage**:
   ```bash
   ./true
   ```

#### wc

1. **Build**:

   ```bash
   go build -o wc cmd/wc/main.go
   ```

2. **Usage**:
   ```bash
   ./wc [flags] <file>...
   ```

#### yes

1. **Build**:

   ```bash
   go build -o yes cmd/yes/main.go
   ```

2. **Usage**:
   ```bash
   ./yes [text]
   ```

### Usage

Execute commands using the built binary for each specific command.

```bash
./<command> [flags] (arg || [<args list>])
```

## Examples

- Print the first 10 lines of a file:

  ```bash
  ./head -n 10 myfile.txt
  ```

- Print the last 10 lines of a file:
  ```bash
  ./tail -n 10 myfile.txt
  ```
