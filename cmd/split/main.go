package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/sc0t2/split/pkg/split"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	os.Exit(run())
}

func run() int {
	filePath := flag.String("file", "", "the csv file to split")
	outDir := flag.String("outDir", ".", "the directory to store the split files in")
	splitSize := flag.Int("size", 1000, "the number of rows to split on")
	flag.Parse()

	file := filepath.Base(*filePath)
	if !strings.HasSuffix(file, "csv") {
		fmt.Printf("file %s does not have csv extension\n", *filePath)
		return 1
	}

	inFile, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("could not open file: %q\n", err)
		return 2
	}

	reader := csv.NewReader(inFile)
	fileCount, err := split.Csv(reader, *splitSize, *outDir, strings.Split(file, ".")[0]+"-")
	if err != nil {
		fmt.Printf("error while splitting file: %q\n", err)
		return 3
	}

	fmt.Printf("Split %s into %d files\n", *filePath, fileCount)
	return 0
}
