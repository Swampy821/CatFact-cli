package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type catFact struct {
	Facts []string
}

func getCatFact() string {
	resp, _ := http.Get("http://catfacts-api.appspot.com/api/facts")
	body, _ := ioutil.ReadAll(resp.Body)
	bodyFix := strings.Replace(string(body), "facts", "Facts", -1)
	return bodyFix
}

func parseIntoObject(fact string) catFact {
	factBytes := []byte(fact)
	var theFact catFact

	json.Unmarshal(factBytes, &theFact)
	return theFact
}

func outputCatFact(fact catFact) {
	fmt.Println("")
	fmt.Println(fact.Facts[0])
	fmt.Println("")
}

func main() {
	count := flag.Int("count", 1, "Number of catfacts to get")
	flag.Parse()

	for i := 0; i < *count; i++ {
		catFact := getCatFact()

		fact := parseIntoObject(catFact)

		outputCatFact(fact)
	}
}
