package main

import (
	"bufio"
	"flag"
	"fmt"
	"gitlab.ozon.ru/test/app"
	"os"
)

var t = flag.String("type", "", "type of source")

const pullSize = 5
const substring = "Go "

func main() {
	flag.Parse()
	sources := bufio.NewScanner(os.Stdin)
	processor, err := app.GetProcessor(*t)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := app.NewRunner(pullSize, substring, processor)
	for sources.Scan() {
		r.Run(sources.Text())
	}
	r.GetResults()
}


