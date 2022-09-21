package messageparser

import (
	"bufio"
	"encoding/csv"
	"os"
)

type CSVParser struct {
}

func (c *CSVParser) Parse(inputFile string) ([][]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(file))
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
