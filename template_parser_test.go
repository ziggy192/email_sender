package email_sender

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const testEmailTemplate = "etc/test/email_template.json"

func TestFileTemplateParser_ParseEmails(t *testing.T) {
	t.Parallel()

	today := time.Now().Format("02 Jan 2006")
	tests := []struct {
		name string
		cs   []*Customer
		want []*Email
	}{
		{
			name: "success",
			cs: []*Customer{
				{
					Title:     "MR",
					FirstName: "John",
					LastName:  "Smith",
					Email:     "john.smith@example.com",
				},
			},
			want: []*Email{
				{
					From:     "The marketing team <marketing@example.com>",
					To:       "john.smith@example.com",
					Subject:  "A new product is being launched...",
					MimeType: "text/plain",
					Body:     "Hi MR John Smith,\nToday, " + today + ", we would like to tell you that... Sincerely,\nThe Marketing Team",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			template, err := NewFileTemplateParser(testEmailTemplate)
			require.NoError(t, err)
			got := template.ParseEmails(tt.cs)
			require.Equal(t, tt.want, got)
		})
	}
}
