# wc.go

`wc.go` is a command-line application written in Go that mimics the functionality of the Unix `wc` command. It counts the number of words, lines, and characters in the given file.

## Options

- `-f`: Specify the file to read from. If this option is not provided, the program will read from stdin.
- `-l`: Print only the number of lines.
- `-w`: Print only the number of words.
- `-c`: Print only the number of characters.
- `-r`: Print only the number of unicode characters (runes).
- `-b`: Print only the number of bytes.
- `-h`: Print help.

## Usage

```bash
# This will print the number of lines, words, and characters in myfile.txt
go run wc.go -f myfile.txt
```

## Installation

To install wc.go, you can clone this repository and build the program using the Go compiler:

```bash
git clone https://github.com/roman91DE/wc
cd yourrepository
go build wc.go
```

This will create an executable file named wc.go in your current directory.

