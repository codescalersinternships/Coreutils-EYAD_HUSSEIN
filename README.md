# Coreutils-EYAD_HUSSEIN

Coreutils-EYAD_HUSSEIN is a collection of Unix-like command line utilities implemented in Go.

## Project Structure

```bash
.
├── bin
├── pkg
├── README.md
└── src
    └── Coreutils-EYAD_HUSSEIN
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
        ├── main
        │   ├── main
        │   └── main.go
        ├── models
        │   ├── command.go
        │   └── commands.go
        └── utils
            ├── generalUtils.go
            └── validateInput.go
```


- **bin/**: Contains executable binaries.
- **pkg/**: Holds package files.
- **src/Coreutils-EYAD_HUSSEIN/**: Source directory for the project.

## Commands Implemented

- **cat**: Concatenate files and print on the standard output.
  - Allowed flags: `-n`

- **echo**: Display line of text in standard output.
  - Allowed flags: `-n`

- **env**: List all environment variables.
  - No additional flags allowed.

- **false**: Do nothing, unsuccessfully..
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

1. **Installation**: Clone the repository and navigate to `src/Coreutils-EYAD_HUSSEIN`.

2. **Build**: Compile the project using `go build`.

    ```bash
    go build -o coreutils main/main.go
    ```

3. **Usage**: Execute commands using the built binary `coreutils`.

    ```bash
    ./coreutils <command> [flags] (arg || [<args list>])
    ```

## Examples

- Print the first 20 lines of a file:
  
  ```bash
  ./coreutils head -n 10 myfile.txt

- Print the last 20 lines of a file:
  
  ```bash
  ./coreutils tail -n 10 myfile.txt
