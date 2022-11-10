package email_sender

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileEmailSender_Send(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		emails      []*Email
		want        []*ErrEmail
		wantErr     bool
		wantContent string
	}{
		{
			name: "success",
			emails: []*Email{
				{
					From:     "someone@gmail.com",
					To:       "someoneelse@gmail.com",
					Subject:  "Hello",
					MimeType: "text/plain",
					Body:     "Hello there my name is someone",
				},
			},
			want: []*ErrEmail{
				{
					Email: &Email{
						From:     "someone@gmail.com",
						To:       "someoneelse@gmail.com",
						Subject:  "Hello",
						MimeType: "text/plain",
						Body:     "Hello there my name is someone",
					},
					Err: nil,
				},
			},
			wantErr: false,
			wantContent: `{
    "from": "someone@gmail.com",
    "to": "someoneelse@gmail.com",
    "subject": "Hello",
    "mime_type": "text/plain",
    "body": "Hello there my name is someone"
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// no parallel
			outDir, err := os.MkdirTemp("", "TestSender_"+tt.name)
			println(outDir)
			require.NoError(t, err)
			defer os.RemoveAll(outDir)

			f, err := NewFileEmailSender(outDir)
			require.NoError(t, err)

			got, err := f.Send(tt.emails)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)

			content, err := os.ReadFile(outDir + "/sent_email_1.json")
			require.NoError(t, err)
			require.Equal(t, tt.wantContent, string(content))
		})
	}
}
