package main

import (
	"fmt"
	"time"

	"github.com/pkg/browser"
)

var (
	baseURL                  = "https://www.bing.com/search?PC=U523&q=%s&FORM=ANAB01"
	searchTerm               = "this+should+be+long+enough+to+work"
	waitTime   time.Duration = 7
)

func main() {

	println("searching")
	for {
		if len(searchTerm) == 0 {
			break
		}

		url := fmt.Sprintf(baseURL, searchTerm)
		println("searching for " + url)
		browser.OpenURL(url)

		searchTerm = searchTerm[:len(searchTerm)-1]
		time.Sleep(waitTime * time.Second)
	}
	println("finished")
}
