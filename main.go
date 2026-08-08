package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define CLI flags
	sourceDir := flag.String("source", "./content", "Directory containing markdown files")
	templateDir := flag.String("template", "./templates", "Directory containing HTML templates")
	outputDir := flag.String("output", "./public", "Directory for the generated static site")

	// Parse the flags
	flag.Parse()

	// Welcome message and current configuration
	fmt.Println("🚀 Starting GoSSG - Static Site Generator")
	fmt.Printf("Source Directory  : %s\n", *sourceDir)
	fmt.Printf("Template Directory: %s\n", *templateDir)
	fmt.Printf("Output Directory  : %s\n", *outputDir)

	// Basic validation: ensure source directory exists
	if _, err := os.Stat(*sourceDir); os.IsNotExist(err) {
		fmt.Printf("Error: Source directory '%s' does not exist.\n", *sourceDir)
		os.Exit(1)
	}

	// Basic validation: ensure template directory exists
	if _, err := os.Stat(*templateDir); os.IsNotExist(err) {
		fmt.Printf("Error: Template directory '%s' does not exist.\n", *templateDir)
		os.Exit(1)
	}

	// TODO: Add file processing and markdown parsing logic here
	fmt.Println("\n✅ Configuration loaded successfully. Ready for the next step!")
}
