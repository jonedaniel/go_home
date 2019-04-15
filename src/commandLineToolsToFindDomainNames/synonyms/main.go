package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"commandLineToolsToFindDomainNames/thesaurus"
)

func main() {
	apiKey := "c1cb5c8283ad9b1c4c156877d9dcc7a6"
	hugh := &thesaurus.BigHugh{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := hugh.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for \""+word+"\"", err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for \"" + word + "\"")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
