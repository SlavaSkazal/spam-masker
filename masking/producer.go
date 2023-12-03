package masking

import (
	"bufio"
	"fmt"
	"os"
)

type producer interface {
	produce() ([]string, error)
}

type Producer struct {
	filepathSource string
}

func (p *Producer) produce() ([]string, error) {
	var maskingStrs []string

	f, err := os.Open(p.filepathSource)
	if err != nil {
		return maskingStrs, fmt.Errorf("error os.Open: %s: %w", p.filepathSource, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		maskingStrs = append(maskingStrs, sc.Text())
	}

	return maskingStrs, nil
}
