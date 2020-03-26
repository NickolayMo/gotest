package processor

import (
	"io/ioutil"
	"os"
)

type FileProcessor struct {
}

func (u *FileProcessor) Process(source string)(string, error) {
	reader, err := os.Open(source)
	if err != nil {
		return "", err
	}
	defer reader.Close()
	text, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(text), err
}