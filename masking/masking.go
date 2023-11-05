package masking

import (
	"bufio"
	"fmt"
	"os"
)

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present([]string) error
}

type Producer struct {
	filepathSource string
}

type Presenter struct {
	filepathRes string
}

type Service struct {
	prod producer
	pres presenter
}

func newProducer(filepathSource string) *Producer {
	return &Producer{filepathSource: filepathSource}
}

func newPresenter(filepathRes string) *Presenter {
	return &Presenter{filepathRes: filepathRes}
}

func NewService(filepathSource string, filepathRes string) *Service {
	var prod producer = newProducer(filepathSource)
	var pres presenter = newPresenter(filepathRes)

	return &Service{prod, pres}
}

func (p *Producer) produce() ([]string, error) {
	var maskingStrs []string

	f, err := os.Open(p.filepathSource)
	if err != nil {
		return maskingStrs, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		maskingStrs = append(maskingStrs, sc.Text())
	}
	return maskingStrs, nil
}

func (p *Presenter) present(maskedStrs []string) error {
	f, err := os.Create(p.filepathRes)
	fmt.Println(p.filepathRes)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range maskedStrs {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) Run() error {
	maskingStrs, err := s.prod.produce()
	if err != nil {
		return err
	}

	maskedStrs := make([]string, len(maskingStrs))
	for i, maskStr := range maskingStrs {
		maskedStr := s.findAndMaskLinks(maskStr)
		maskedStrs[i] = maskedStr
	}

	err = s.pres.present(maskedStrs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) findAndMaskLinks(sourceStr string) string {
	var mask = []rune("http://")

	if len(sourceStr) < len(mask) {
		return sourceStr
	}

	maskingStr, maskedStr := make([]rune, len([]rune(sourceStr))), make([]rune, len([]rune(sourceStr)))
	copy(maskingStr, []rune(sourceStr))
	copy(maskedStr, []rune(sourceStr))

	space := []rune(" ")[0]
	star := []rune("*")[0]
	var itsLink bool
	var indBeginLink int

	for i := 0; i < len(maskingStr); i++ {
		if !itsLink {
			if i > len(maskingStr)-len(mask) {
				break
			}

			if string(maskingStr[i:i+len(mask)]) == string(mask) && (i == 0 || maskingStr[i-1] == space) {
				itsLink = true
				i += len(mask)
				indBeginLink = i
			}
		} else if itsLink {
			if maskingStr[i] == space || i+1 >= len(maskingStr) {
				lenStars := i - indBeginLink

				if i+1 >= len(maskingStr) {
					lenStars++
				}

				for k := 0; k < lenStars; k++ {
					maskedStr[indBeginLink] = star
					indBeginLink++
				}
				itsLink = false
			}
		}
	}

	return string(maskedStr)
}
