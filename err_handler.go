package email_sender

import "log"

type ErrExporter struct {
}

func NewErrExporter() *ErrExporter {
	return &ErrExporter{}
}

func (e ErrExporter) HandleErr(emails []*ErrEmail) error {
	if len(emails) == 0 {
		return nil
	}
	// todo
	log.Printf("hanlding errors ")
	return nil
}
