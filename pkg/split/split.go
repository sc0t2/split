package split

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Csv splits the file at filePath into smaller CSV files with chunkSize rows, returning the number of files created
func Csv(reader *csv.Reader, chunkSize int, outDir, prefix string) (int, error) {
	header, err := reader.Read()
	if err != nil {
		return 0, err
	}

	fileCount := 0

	for {
		chunk, chunkErr := getChunk(reader, chunkSize)

		if chunkErr != nil && chunkErr != io.EOF {
			return fileCount, chunkErr
		}

		// do not write empty chunks
		emptyRow := false
		if len(chunk) == 1 {
			emptyRow = true
			for _, cell := range chunk[0] {
				if len(cell) > 0 {
					emptyRow = false
					break
				}
			}
		}

		if !emptyRow {
			outFile, createErr := os.Create(fmt.Sprintf("%s/%s%d.csv", outDir, prefix, fileCount))
			if createErr != nil {
				return 0, createErr
			}

			writer := csv.NewWriter(outFile)
			writeErr := writer.Write(header)
			if writeErr != nil {
				return fileCount, writeErr
			}
			writeErr = writer.WriteAll(chunk)
			if writeErr != nil {
				return fileCount, writeErr
			}
			writer.Flush()

			fileCount++
			closeErr := outFile.Close()
			if closeErr != nil {
				return fileCount, closeErr
			}
		}

		if chunkErr == io.EOF {
			break
		}
	}

	return fileCount, nil
}

// get the chunk of the CSV with size chunkSize
func getChunk(r *csv.Reader, chunkSize int) ([][]string, error) {
	var records [][]string
	for i := 0; i < chunkSize; i++ {
		record, readErr := r.Read()
		if readErr == io.EOF {
			records = append(records, record)
			return records, io.EOF
		}
		if readErr != nil {
			return nil, readErr
		}
		records = append(records, record)
	}
	return records, nil
}
