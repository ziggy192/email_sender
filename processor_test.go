package email_sender

import (
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestEmailProcessor_Process(t *testing.T) {
	t.Parallel()

	type fields struct {
		reader         *MockReader
		sender         *MockSender
		errHandler     *MockErrHandler
		templateParser *MockTemplateParser
	}

	tests := []struct {
		name         string
		prepareMocks func(t *testing.T, fields fields)
		n            int
		want         bool
		wantErr      bool
	}{
		{
			name: "success",
			prepareMocks: func(t *testing.T, f fields) {
				var mockCustomers = make([]*Customer, 0)
				for i := 0; i < 3; i++ {
					mockCustomers = append(mockCustomers, &Customer{
						Title:     RandomString(10, AlphaNumericCharacters),
						FirstName: RandomString(10, AlphaNumericCharacters),
						LastName:  RandomString(10, AlphaNumericCharacters),
						Email:     RandomString(10, AlphaNumericCharacters),
					})
				}
				f.reader.EXPECT().Read(10).Return(mockCustomers, io.EOF)

				mockEmails := make([]*Email, 0)
				for i := 0; i < 3; i++ {
					mockEmails = append(mockEmails, &Email{
						From:     RandomString(10, AlphaNumericCharacters),
						To:       RandomString(10, AlphaNumericCharacters),
						Subject:  RandomString(10, AlphaNumericCharacters),
						MimeType: "plain/text",
						Body:     RandomString(10, AlphaNumericCharacters),
					})
				}
				f.templateParser.EXPECT().ParseEmails(gomock.Any()).DoAndReturn(
					func(cs []*Customer) []*Email {
						require.Len(t, cs, 3)
						return mockEmails
					})

				f.sender.EXPECT().Send(mockEmails).Return(nil, nil)
				f.errHandler.EXPECT().HandleErr(gomock.Len(0))
			},
			n:       10,
			want:    false,
			wantErr: false,
		},
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := fields{
				reader:         NewMockReader(ctrl),
				sender:         NewMockSender(ctrl),
				errHandler:     NewMockErrHandler(ctrl),
				templateParser: NewMockTemplateParser(ctrl),
			}

			tt.prepareMocks(t, f)

			e := &EmailProcessor{
				reader:         f.reader,
				sender:         f.sender,
				errHandler:     f.errHandler,
				templateParser: f.templateParser,
			}
			got, err := e.Process(tt.n)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
