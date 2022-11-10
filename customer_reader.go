package email_sender

import (
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/jszwec/csvutil"
)

// CustomerReader reads customers from csv file
type CustomerReader struct {
	customerFile *os.File
	dec          *csvutil.Decoder
}

func NewCustomerReader(customerPath string) (er *CustomerReader, err error) {
	var customerFile *os.File
	cleanUp := func() {
		if customerFile != nil {
			_ = customerFile.Close()
		}
	}

	defer func() {
		if err != nil {
			cleanUp()
		}
	}()

	customerFile, err = os.Open(customerPath)
	if err != nil {
		LogErr(err)
		return
	}

	dec, err := csvutil.NewDecoder(csv.NewReader(customerFile))
	if err != nil {
		LogErr(err)
		return
	}
	return &CustomerReader{
		customerFile: customerFile,
		dec:          dec,
	}, nil
}

// Read reads upto n next customers. Return io.EOF when no next customers to read
func (cr *CustomerReader) Read(n int) ([]*Customer, error) {
	customers := make([]*Customer, 0, n)
	for i := 0; i < n; i++ {
		var c *Customer
		err := cr.dec.Decode(&c)
		if err != nil && err != io.EOF {
			LogErr(err)
			return nil, err
		} else if errors.Is(err, io.EOF) {
			return customers, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (cr *CustomerReader) Close() {
	if cr.customerFile != nil {
		_ = cr.customerFile.Close()
		cr.customerFile = nil
	}

}
