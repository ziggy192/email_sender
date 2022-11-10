package email_sender

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

// ErrExporter export error customers to file
type ErrExporter struct {
	encoder *csvutil.Encoder
	errFile *os.File
	writer  *csv.Writer
}

func NewErrExporter(errPath string) (*ErrExporter, error) {
	_, err := os.Stat(errPath)
	if err == nil {
		err = fmt.Errorf("error file %s already exists", errPath)
		LogErr(err)
		return nil, err
	}

	file, err := os.Create(errPath)
	if err != nil {
		LogErr(err)
		return nil, err
	}

	writer := csv.NewWriter(file)
	encoder := csvutil.NewEncoder(writer)
	if err != nil {
		LogErr(err)
		return nil, err
	}

	return &ErrExporter{
		encoder: encoder,
		errFile: file,
		writer:  writer,
	}, nil
}

func (e *ErrExporter) Close() {
	_ = e.errFile.Close()
}

func (e *ErrExporter) HandleErr(customers []*Customer) error {
	if len(customers) == 0 {
		return nil
	}

	for _, c := range customers {
		err := e.encoder.Encode(c)
		if err != nil {
			LogErr(err)
			return err
		}
	}

	e.writer.Flush()
	if err := e.writer.Error(); err != nil {
		LogErr(err)
		return err
	}

	return nil
}
