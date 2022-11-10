package email_sender

import (
	"errors"
	"io"
)

type EmailProcessor struct {
	reader     Reader
	sender     Sender
	errHandler ErrHandler
}

func NewEmailProcessor(reader Reader, sender Sender, errHandler ErrHandler) *EmailProcessor {
	return &EmailProcessor{reader: reader, sender: sender, errHandler: errHandler}
}

func (e *EmailProcessor) Process(n int) (bool, error) {
	var next = true
	emails, err := e.reader.Read(n)
	if errors.Is(err, io.EOF) {
		next = false
	} else if err != nil {
		return next, err
	}

	sentEmails, err := e.sender.Send(emails)
	if err != nil {
		return next, err
	}
	errs := make([]*ErrEmail, 0)
	for _, s := range sentEmails {
		if s.Err != nil {
			errs = append(errs, s)
		}
	}
	return next, e.errHandler.HandleErr(errs)
}

type Reader interface {
	Read(n int) ([]*Email, error)
}

type Sender interface {
	Send(emails []*Email) ([]*ErrEmail, error)
}

type ErrHandler interface {
	HandleErr(emails []*ErrEmail) error
}
