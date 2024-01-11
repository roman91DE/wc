package main

// wc prints the number of words, lines and characters in the given file.
// options:
//   -f: file to read from, if empty read from stdin
//   -l: print only the number of lines
//   -w: print only the number of words
//   -c: print only the number of characters
//   -r: print only the number of unicode characters (runes)
// 	 -b: print only the number of bytes
//   -h: print help

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	var (
		args struct {
			file  string
			lines bool
			words bool
			chars bool
			runes bool
			bytes bool
		}

		counter struct {
			lines int
			words int
			chars int
			runes int
			bytes int
		}
		err  error
		mode string
		file *os.File
	)

	flag.StringVar(&args.file, "f", "", "file to read from, if empty read from stdin")
	flag.BoolVar(&args.lines, "l", false, "print only the number of lines")
	flag.BoolVar(&args.words, "w", false, "print only the number of words")
	flag.BoolVar(&args.chars, "c", false, "print only the number of characters")
	flag.BoolVar(&args.runes, "r", false, "print the number of unicode characters (runes)")
	flag.BoolVar(&args.bytes, "b", false, "print the number of bytes")

	flag.Parse()

	flagCount := 0
	if args.lines {
		mode = "l"
		flagCount++
	}
	if args.words {
		mode = "w"
		flagCount++
	}
	if args.chars {
		mode = "c"
		flagCount++
	}
	if args.runes {
		mode = "r"
		flagCount++
	}
	if args.bytes {
		mode = "b"
		flagCount++
	}

	if flagCount > 1 {
		fmt.Println("Error: Only one of -l, -w, -c, -r, or -b can be specified.")
		os.Exit(1)
	}

	if args.file == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(args.file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		switch mode {
		case "b":
			counter.bytes += len([]byte(line))
		case "r":
			counter.runes += len([]rune(line))
		case "l":
			counter.lines++
		case "w":
			counter.words += len(strings.Fields(line))
		case "c":
			counter.chars += len(line)
		default:
			counter.lines++
			counter.words += len(strings.Fields(line))
			counter.chars += len(line)
		}
	}

	switch mode {
	case "l":
		fmt.Println(counter.lines)
	case "w":
		fmt.Println(counter.words)
	case "c":
		fmt.Println(counter.chars)
	case "r":
		fmt.Println(counter.runes)
	case "b":
		fmt.Println(counter.bytes)
	default:
		fmt.Printf("%d %d %d\n", counter.lines, counter.words, counter.chars)
	}
}
