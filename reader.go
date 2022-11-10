package email_sender

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"

	"github.com/jszwec/csvutil"
)

type EmailReader struct {
	template     *EmailTemplate
	templateFile *os.File
	customerFile *os.File
	dec          *csvutil.Decoder
}

func NewEmailReader(templateFilePath, customerPath string) (er *EmailReader, err error) {
	var templateFile, customerFile *os.File
	cleanUp := func() {
		if customerFile != nil {
			_ = customerFile.Close()
		}
		if templateFile != nil {
			_ = templateFile.Close()
		}
	}

	defer func() {
		if err != nil {
			cleanUp()
		}
	}()

	templateFile, err = os.Open(templateFilePath)
	if err != nil {
		LogErr(err)
		return nil, err
	}

	var template = new(EmailTemplate)
	if err := json.NewDecoder(templateFile).Decode(&template); err != nil {
		LogErr(err)
		return nil, err
	}

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
	return &EmailReader{
		template:     template,
		templateFile: templateFile,
		customerFile: customerFile,
		dec:          dec,
	}, nil
}

func (er *EmailReader) Read(n int) ([]*Email, error) {
	emails := make([]*Email, 0, n)
	for i := 0; i < n; i++ {
		var c *Customer
		if err := er.dec.Decode(&c); err == io.EOF {
			emails = append(emails, er.parseEmail(c))
			return emails, io.EOF
		} else if err != nil && err != io.EOF {
			LogErr(err)
			return nil, err
		}
		emails = append(emails, er.parseEmail(c))
	}
	return emails, nil
}

func (er *EmailReader) parseEmail(c *Customer) *Email {
	today := time.Now().Format("02 Jan 2006")
	r := strings.NewReplacer("{{TITLE}}", c.Title, "{{FIRST_NAME}}", c.FirstName, "{{LAST_NAME}}", c.LastName, "{{TODAY}}", today)
	body := r.Replace(er.template.Body)
	return &Email{
		From:     er.template.From,
		To:       c.Email,
		Subject:  er.template.Subject,
		MimeType: er.template.MimeType,
		Body:     body,
	}
}
func (er *EmailReader) Close() {
	if er.customerFile != nil {
		_ = er.customerFile.Close()
		er.customerFile = nil
	}
	if er.templateFile != nil {
		_ = er.templateFile.Close()
		er.templateFile = nil
	}
}
