package service

import (
	"context"
	"encoding/csv"
	"io"
	"os"

	"testtask/model"
)

type (
	// Reader interface for reading data from source
	Reader interface {
		Read(ctx context.Context) error
	}
	// CSVReader implements method of Reader interface for reading csv file.
	CSVReader struct {
		fileName  string
		delimiter rune
		outChan   chan *model.BasicTitle
	}
)

// NewCSVReader returns new instance of CSVReader.
func NewCSVReader(filename string, outChan chan *model.BasicTitle) Reader {
	return &CSVReader{
		fileName:  filename,
		outChan:   outChan,
		delimiter: '\t',
	}
}

// Read reads data from csv file.
func (r *CSVReader) Read(ctx context.Context) error {
	file, err := os.Open(r.fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	defer close(r.outChan)

	rd := csv.NewReader(file)
	rd.Comma = r.delimiter
	for {
		row, err := rd.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		if len(row) < 9 {
			continue
		}
		r.outChan <- &model.BasicTitle{
			Tconst:         row[0],
			TitleType:      row[1],
			PrimaryTitle:   row[2],
			OriginalTitle:  row[3],
			IsAdult:        row[4],
			StartYear:      row[5],
			EndYear:        row[6],
			RuntimeMinutes: row[7],
			Genres:         row[8],
		}

		select {
		case <-ctx.Done():
			return nil
		default:

		}

	}

	return nil
}
