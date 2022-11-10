package email_sender

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrExporter_HandleErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		customers   []*Customer
		wantErr     bool
		wantContent string
	}{{
		name: "success",
		customers: []*Customer{
			{
				Title:     "MR",
				FirstName: "John",
				LastName:  "Smith",
				Email:     "john.smith@example.com",
			},
			{
				Title:     "Mrs",
				FirstName: "Michelle",
				LastName:  "Smith",
				Email:     "michelle.smith@example.com",
			},
		},
		wantErr:     false,
		wantContent: "TITLE,FIRST_NAME,LAST_NAME,EMAIL\nMR,John,Smith,john.smith@example.com\nMrs,Michelle,Smith,michelle.smith@example.com\n",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// no parallel
			errFile := RandomString(10, AlphaNumericCharacters) + "_errors.csv"
			e, err := NewErrExporter(errFile)
			require.NoError(t, err)

			err = e.HandleErr(tt.customers)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			content, err := os.ReadFile(errFile)
			require.NoError(t, err)
			defer os.Remove(errFile)
			require.Equal(t, tt.wantContent, string(content))
		})
	}
}
