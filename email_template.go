package email_sender

type EmailTemplate struct {
	From     string `json:"from"`
	Subject  string `json:"subject"`
	MimeType string `json:"mimeType"`
	Body     string `json:"body"`
}
