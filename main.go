package main

import (
	"fmt"
	"os"
	"spamMasker/masking"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("order numbers not sent")
		return
	}

	var filepathRes string
	filepathSource := os.Args[1]

	if len(os.Args) > 2 {
		filepathRes = os.Args[2]
	} else {
		filepathRes = "output.txt"
	}

	serv := masking.NewService(filepathSource, filepathRes)

	err := serv.Run()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done.")
	}
}
