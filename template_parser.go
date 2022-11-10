package email_sender

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type TemplateParser struct {
	template     *EmailTemplate
	templateFile *os.File
}

func NewTemplateParser(templateFilePath string) (tp *TemplateParser, err error) {
	var templateFile *os.File
	defer func() {
		if err != nil && templateFile != nil {
			_ = templateFile.Close()
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

	return &TemplateParser{
		template:     template,
		templateFile: templateFile,
	}, nil
}

func (t *TemplateParser) ParseEmails(cs []*Customer) []*Email {
	var emails = make([]*Email, 0, len(cs))
	for _, c := range cs {
		today := time.Now().Format("02 Jan 2006")
		r := strings.NewReplacer("{{TITLE}}", c.Title, "{{FIRST_NAME}}", c.FirstName, "{{LAST_NAME}}", c.LastName, "{{TODAY}}", today)
		body := r.Replace(t.template.Body)
		emails = append(emails, &Email{
			From:     t.template.From,
			To:       c.Email,
			Subject:  t.template.Subject,
			MimeType: t.template.MimeType,
			Body:     body,
		})
	}
	return emails
}

func (t *TemplateParser) Close() {
	if t.templateFile != nil {
		_ = t.templateFile.Close()
		t.templateFile = nil
	}
}
