package masking

import (
	"fmt"
	"os"
)

type presenter interface {
	present([]string) error
}

type Presenter struct {
	filepathRes string
}

func (p *Presenter) present(maskedStrs []string) error {
	f, err := os.Create(p.filepathRes)

	if err != nil {
		return fmt.Errorf("error os.Create: %s: %w", p.filepathRes, err)
	}
	defer f.Close()

	for _, line := range maskedStrs {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error WriteString in line %s: file: %s: %w", line, p.filepathRes, err)
		}
	}

	return nil
}
