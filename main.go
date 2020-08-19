package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("ERROR: please specify a file path as an argument")
	}

	// read the current content as a string
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	// replace the first "date: ..." line with the current timestamp
	re := regexp.MustCompile("(?s)^(.*?\n)date: [^\n]*(\n.*)$")
	newText := re.ReplaceAllString(text, fmt.Sprintf("${1}date: %s$2", time.Now().Format("2006-01-02T15:04:05-0700")))

	// overwrite the file with the new content
	err = ioutil.WriteFile(os.Args[1], []byte(newText), 0)
	if err != nil {
		log.Fatal(err)
	}
}
