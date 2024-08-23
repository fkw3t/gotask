package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func ParseDate(date string) (*time.Time, error) {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	t, err := time.ParseInLocation(TimeFormat, date, location)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %v\n", err)
	}

	return &t, nil
}

func ParseNullableDate(date string) (*time.Time, error) {
	if date == "" {
		return nil, nil
	}

	return ParseDate(date)
}

func FormatNullableDate(date *time.Time) string {
	if date == nil {
		return ""
	}

	return date.Format(TimeFormat)
}

func ReadCSV(filepath, filename string) ([][]string, error) {
	file, err := os.Open(fmt.Sprintf("%v/%v.csv", filepath, filename))
	if err != nil {
		return nil, fmt.Errorf("failed to open storage file: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read storage file: %v\n", err)
	}

	return records, nil
}

func WriteCSV(filepath, filename string, content [][]string) error {
	file, err := os.Create(fmt.Sprintf("%v/%v.csv", filepath, filename))
	if err != nil {
		return fmt.Errorf("failed to open storage file: %v\n", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(content)
	if err != nil {
		return fmt.Errorf("failed to write to storage file: %v\n", err)
	}

	return nil
}

func AppendCSV(filepath, filename string, content []string) error {
	file, err := os.OpenFile(fmt.Sprintf("%v/%v.csv", filepath, filename), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open storage file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(content); err != nil {
		return fmt.Errorf("failed to write record to file: %v", err)
	}
	return nil
}
