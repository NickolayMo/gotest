package app

import (
	"fmt"
	"gitlab.ozon.ru/test/app/processor"
	"strings"
)

func GetProcessor(t string) (Processor, error) {
	switch t {
	case "url":
		return &processor.UrlProcessor{}, nil
	case "file":
		return &processor.FileProcessor{}, nil
	default:
		return nil, fmt.Errorf("unexpected type %s", t)
	}
}

func Count(text, substring string) int {
	return strings.Count(text, substring)

}
