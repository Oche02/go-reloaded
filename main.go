package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("ATTENTION: expexting>> go run . sample.txt result.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("ATTENTION: failed to open input file", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("ATTENTION: failed to create output file", err)
		return
	}
	defer outFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("ATTENTION: error detected")
		return
	}

	writer := bufio.NewWriter(outFile)
	for idx, line := range lines {
		result := processFile(line)
		if idx < len(lines)-1 {
			fmt.Fprintln(writer, result)
		} else {
			fmt.Fprint(writer, result)
		}
	}
	writer.Flush()
	fmt.Println("Successfil")
}
