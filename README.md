# shell-go

A simple shell implementation in Go, designed to execute commands and change directories.

## Features

- Execute system commands.
- Change the current working directory.
- Exit the shell.

## Installation

To install and run `shell-go`, you need to have Go installed on your system. If you haven't installed Go yet, you can download it from the [official Go website](https://golang.org/dl/).

1. Clone the repository or download the source code.
2. Open a terminal and navigate to the directory containing the `app.go` file.
3. Run the following command to build the executable:

```bash
go build -o shell-go app.go
```

4. Execute the shell by running:

```bash
./shell-go
```

## Usage

After starting `shell-go`, you can type commands directly into the terminal. Here are some examples of commands you can use:

- `cd <directory>`: Change the current working directory to `<directory>`.
- `exit`: Exit the shell.
- Any other command: Execute the command as if you were in a regular shell.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on the GitHub repository. If you'd like to contribute code, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
