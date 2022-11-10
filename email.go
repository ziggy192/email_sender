package email_sender

type Email struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	MimeType string `json:"mime_type"`
	Body     string `json:"body"`
}
