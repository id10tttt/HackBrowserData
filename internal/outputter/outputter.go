package outputter

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	jsoniter "github.com/json-iterator/go"

	"hack-browser-data/internal/browingdata"
)

type outPutter struct {
	json bool
	csv  bool
}

func NewOutPutter(flag string) *outPutter {
	o := &outPutter{}
	if flag == "json" {
		o.json = true
	} else {
		o.csv = true
	}
	return o
}

func (o *outPutter) Write(data browingdata.Source, writer io.Writer) error {
	switch o.json {
	case true:
		encoder := jsoniter.NewEncoder(writer)
		encoder.SetIndent("  ", "  ")
		encoder.SetEscapeHTML(false)
		return encoder.Encode(data)
	default:
		gocsv.SetCSVWriter(func(w io.Writer) *gocsv.SafeCSVWriter {
			writer := csv.NewWriter(w)
			writer.Comma = ','
			return gocsv.NewSafeCSVWriter(writer)
		})
		return gocsv.Marshal(data, writer)
	}
}

func (o *outPutter) CreateFile(dir, filename string) (*os.File, error) {
	if filename == "" {
		return nil, errors.New("empty filename")
	}

	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0777)
			if err != nil {
				return nil, err
			}
		}
	}

	var file *os.File
	var err error
	p := filepath.Join(dir, filename)
	file, err = os.OpenFile(p, os.O_TRUNC|os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}
