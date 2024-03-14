package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "", "Path to the PEM file")
	directory := flag.String("dir", "", "Direcory containing PEM files")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Filename required")
		os.Exit(1)
	}

	command := flag.Args()[0]

	switch command {
	case "info":
		switch {
		case *filename != "":
			fileInfo(*filename)
		case *directory != "":
			dirInfo(*directory)
		default:
			fmt.Println("Filename or directory required")
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}

	if *filename == "" {
		fmt.Println("Filename required")
		os.Exit(1)
	}

}
