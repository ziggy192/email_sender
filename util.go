package email_sender

import "log"

func LogErr(err error) {
	log.Printf("[error] %s", err.Error())
}
