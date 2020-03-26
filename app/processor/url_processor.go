package processor

import (
	"io/ioutil"
	"net/http"
)

type UrlProcessor struct {
}

func (u *UrlProcessor) Process(source string) (string, error) {
	result, err := http.Get(source)
	if err != nil {
		return "", err
	}
	defer result.Body.Close()
	text, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}
	return string(text), nil
}
