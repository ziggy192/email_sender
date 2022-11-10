package main

import (
	"flag"

	es "github.com/ziggy192/email_sender"
)

func main() {
	templateFilePath := flag.String("template", "", "(required) path to email template json file")
	customerFilePath := flag.String("customers", "", "(required) path to customers csv file")
	outputDir := flag.String("out", "", "(required) path to output emails directory") // todo create dir if not exists
	errorFilePath := flag.String("error", "", "(required) path to errors csv file")   // todo create if not exists

	// todo make default value for error file
	flag.Parse()

	if len(*templateFilePath) == 0 || len(*customerFilePath) == 0 || len(*outputDir) == 0 || len(*errorFilePath) == 0 {
		flag.PrintDefaults()
		return
	}

	er, err := es.NewEmailReader(*templateFilePath, *customerFilePath)
	if err != nil {
		panic(err)
	}
	defer er.Close()

	sender := es.NewFileEmailSender(*outputDir)

	errHandler := es.NewErrExporter()
	processor := es.NewEmailProcessor(er, sender, errHandler) // todo add errHandler
	for next := true; next; {
		next, err = processor.Process(5) // todo config this ?
		if err != nil {
			panic(err)
		}
	}
}
