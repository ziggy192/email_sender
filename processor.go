package email_sender

import (
	"errors"
	"io"
)

type EmailProcessor struct {
	reader         Reader
	sender         Sender
	errHandler     ErrHandler
	templateParser *TemplateParser
}

func NewEmailProcessor(reader Reader, sender Sender, errHandler ErrHandler, templateParser *TemplateParser) *EmailProcessor {
	return &EmailProcessor{reader: reader, sender: sender, errHandler: errHandler, templateParser: templateParser}
}

func (e *EmailProcessor) Process(n int) (bool, error) {
	var next = true
	customers, err := e.reader.Read(n)
	if errors.Is(err, io.EOF) {
		next = false
	} else if err != nil {
		return next, err
	}

	validated := make([]*Customer, 0, n)
	invalided := make([]*Customer, 0)
	for _, customer := range customers {
		err := customer.Validate()
		if err != nil {
			invalided = append(invalided, customer)
			continue
		}
		validated = append(validated, customer)
	}

	emails := e.templateParser.ParseEmails(validated)
	_, err = e.sender.Send(emails)
	if err != nil {
		return next, err
	}

	err = e.errHandler.HandleErr(invalided)
	if err != nil {
		return next, err
	}

	return next, nil
}

func (c *Customer) Validate() error {
	if len(c.Email) == 0 {
		return errors.New("no email information")
	}
	return nil
}

type Reader interface {
	Read(n int) ([]*Customer, error)
}

type Sender interface {
	Send(emails []*Email) ([]*ErrEmail, error)
}

type ErrHandler interface {
	HandleErr(customers []*Customer) error
}
