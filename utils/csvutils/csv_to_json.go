package csvutils

import (
	"encoding/csv"
	"os"
)

func CSVToJSON(csvFile string) ([]map[string]string, error) {
	f, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	headers, _ := r.Read()
	records, _ := r.ReadAll()

	var result []map[string]string
	for _, row := range records {
		entry := make(map[string]string)
		for i, val := range row {
			entry[headers[i]] = val
		}
		result = append(result, entry)
	}
	return result, nil
}
