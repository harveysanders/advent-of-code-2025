package main

import (
	"fmt"
	"log"
	"os"

	"github.com/harveysanders/aoc2025/day07"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input-file-path>")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer file.Close()

	diagram, err := day07.ParseDiagram(file)
	if err != nil {
		log.Fatalf("Failed to parse diagram: %v", err)
	}

	timelines := diagram.CountTimelines()
	fmt.Printf("Total timelines: %d\n", timelines)
}
