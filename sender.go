package email_sender

import "log"

type ErrEmail struct {
	Email
	Err error `json:"-"`
}

type FileEmailSender struct {
	outDir string
}

func NewFileEmailSender(outDir string) *FileEmailSender {
	return &FileEmailSender{outDir: outDir}
}

// Send sends emails in batch and returns errors in the input order
func (f *FileEmailSender) Send(emails []*Email) ([]*ErrEmail, error) {
	// todo create dir if not exists
	// todo implement this
	for _, email := range emails {
		log.Printf("send email %+v", email)
	}
	return nil, nil
}
