package email_sender

import (
	"fmt"
	"log"
)

func LogErr(v ...any) {
	args := make([]any, 0, 1+len(v))
	args = append(append(args, "[error] "), v...)
	_ = log.Output(2, fmt.Sprint(args...))
}
