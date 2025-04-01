package csvutils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func JSONToCSV(jsonData []map[string]interface{}, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if len(jsonData) == 0 {
		return fmt.Errorf("no data")
	}

	// Write header
	headers := []string{}
	for k := range jsonData[0] {
		headers = append(headers, k)
	}
	writer.Write(headers)

	// Write rows
	for _, row := range jsonData {
		record := []string{}
		for _, h := range headers {
			record = append(record, fmt.Sprintf("%v", row[h]))
		}
		writer.Write(record)
	}

	return nil
}
