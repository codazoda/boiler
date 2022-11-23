package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	templateDir := "~/broiler/template/"
	// If it's set, use an environment variable as the template directory
	if os.Getenv("BOILER") != "" {
		templateDir = os.Getenv("BOILER")
	}
	var templateFile string
	// Make sure there is only one command line argument
	argLength := len(os.Args[1:])
	if argLength != 1 {
		showSyntax()
	}
	// Grab the first command line argument
	template := os.Args[1]
	// Check if a main file exists for this template name
	templateFile = templateDir + template + "/main.txt"
	_, err := os.Stat(templateFile)
	if os.IsNotExist(err) {
		// Check if a matching file exists for this template name
		templateFile = templateDir + template + ".txt"
		_, err := os.Stat(templateFile)
		if os.IsNotExist(err) {
			fmt.Printf("Error: Template file not found (%s)\n", templateFile)
			os.Exit(1) // TODO: Fix exit code
		}
	}
	// TODO: Output the file
	cat(templateFile)
}

// Cat a file to stdout
func cat(fname string) {
	fh, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, fh)
	if err != nil {
		log.Fatal(err)
	}
}

// Output the program syntax
func showSyntax() {
	fmt.Println("Broiler v1.0.0")
	fmt.Println("Syntax: boiler [template]")
	os.Exit(1)
}
