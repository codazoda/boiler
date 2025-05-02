package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	templateDir := ""
	// If it's set, use an environment variable as the template directory
	if os.Getenv("BOILER") != "" {
		templateDir = os.Getenv("BOILER")
	} else {
		fmt.Println("Error: Set the BOILER environment variable to your template path.")
		os.Exit(1)
	}
	var templateFile string
	// Make sure there is only one command line argument
	argLength := len(os.Args[1:])
	if argLength != 1 {
		showSyntax(templateDir)
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
			os.Exit(1)
		}
	}
	// Output the file
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
func showSyntax(templateDir string) {
	fmt.Println("Boiler v1.2")
	fmt.Println("Syntax: boiler [template]")
	os.Exit(1)
}
