package excel

import "github.com/xuri/excelize/v2"

func ReadExcelRows(filePath, sheet string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
