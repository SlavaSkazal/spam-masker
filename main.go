package main

import (
	"log"
	"os"
	"spamMasker/masking"
)

func main() {
	minNumArgs := 2
	numArgPathSource := 1
	numArgPathRes := 2

	if len(os.Args) < minNumArgs {
		log.Println("order numbers not sent")
		return
	}

	var filepathSource, filepathRes string

	filepathSource = os.Args[numArgPathSource]
	if len(os.Args) > minNumArgs {
		filepathRes = os.Args[numArgPathRes]
	} else {
		filepathRes = "data/masking/output.txt"
	}

	serv := masking.NewService(filepathSource, filepathRes)

	if err := serv.Run(); err != nil {
		log.Println(err)
	} else {
		log.Println("Done.")
	}
}
