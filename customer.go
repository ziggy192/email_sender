package email_sender

import "errors"

// Customer defines model for customer
type Customer struct {
	Title     string `csv:"TITLE"`
	FirstName string `csv:"FIRST_NAME"`
	LastName  string `csv:"LAST_NAME"`
	Email     string `csv:"EMAIL"`
}

func (c *Customer) Validate() error {
	if len(c.Email) == 0 {
		return errors.New("no email information")
	}
	return nil
}
