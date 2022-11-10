package main

import (
	"flag"
	"log"

	es "github.com/ziggy192/email_sender"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	templateFilePath := flag.String("template", "", "(required) path to email template json file")
	customerFilePath := flag.String("customers", "", "(required) path to customers csv file")
	outputDir := flag.String("out", "", "(required) path to output emails directory")
	errorFilePath := flag.String("error", "", "(required) path to errors csv file")

	flag.Parse()

	if len(*templateFilePath) == 0 || len(*customerFilePath) == 0 || len(*outputDir) == 0 || len(*errorFilePath) == 0 {
		flag.PrintDefaults()
		return
	}

	er, err := es.NewCustomerReader(*customerFilePath)
	if err != nil {
		panic(err)
	}
	defer er.Close()

	tp, err := es.NewFileTemplateParser(*templateFilePath)
	if err != nil {
		panic(err)
	}
	defer tp.Close()

	sender, err := es.NewFileEmailSender(*outputDir)
	if err != nil {
		panic(err)
	}

	errHandler, err := es.NewErrExporter(*errorFilePath)
	if err != nil {
		panic(err)
	}
	defer errHandler.Close()

	processor := es.NewEmailProcessor(er, sender, errHandler, tp)
	for next := true; next; {
		next, err = processor.Process(5) // todo config this number?
		if err != nil {
			panic(err)
		}
	}
}
