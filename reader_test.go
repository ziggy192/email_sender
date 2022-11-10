package email_sender

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

const customerTestFile = "etc/test/customers.csv"

func TestCustomerReader_Read(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		n       int
		want    []*Customer
		wantErr error
	}{
		{
			name: "success",
			n:    10,
			want: []*Customer{
				{
					Title:     "Mr",
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
				{
					Title:     "Mr",
					FirstName: "something",
					LastName:  "",
					Email:     "",
				},
			},
			wantErr: io.EOF,
		},
		{
			name: "read_less_than_total",
			n:    1,
			want: []*Customer{
				{
					Title:     "Mr",
					FirstName: "John",
					LastName:  "Smith",
					Email:     "john.smith@example.com",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cr, err := NewCustomerReader(customerTestFile)
			require.NoError(t, err)

			got, err := cr.Read(tt.n)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}
