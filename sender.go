package email_sender

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"go.uber.org/atomic"
)

// FileEmailSender send emails to provided file
type FileEmailSender struct {
	outDir  string
	fileIdx *atomic.Int64
}

func NewFileEmailSender(outDir string) (*FileEmailSender, error) {
	entries, err := os.ReadDir(outDir)
	if errors.Is(err, os.ErrNotExist) {
		LogErr("directory ", outDir, " not exists: ", err)
		return nil, err
	}
	return &FileEmailSender{
		outDir:  outDir,
		fileIdx: atomic.NewInt64(int64(len(entries))),
	}, nil
}

// Send sends emails in batch and returns errors in the input order
func (f *FileEmailSender) Send(emails []*Email) ([]*ErrEmail, error) {
	res := make([]*ErrEmail, len(emails))
	for i, email := range emails {
		idx := f.fileIdx.Inc()
		f, err := os.Create(f.outDir + fmt.Sprintf("/sent_email_%d.json", idx))
		if err != nil {
			LogErr(err)
			res[i] = &ErrEmail{Email: email, Err: err}
			continue
		}
		ec := json.NewEncoder(f)
		ec.SetIndent("", "    ")
		ec.SetEscapeHTML(false)
		err = ec.Encode(email)
		if err != nil {
			LogErr(err)
			res[i] = &ErrEmail{Email: email, Err: err}
			continue
		}
		res[i] = &ErrEmail{Email: email}
		_ = f.Close()
	}
	return res, nil
}

type ErrEmail struct {
	*Email
	Err error
}
