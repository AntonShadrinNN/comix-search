package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// getStopWords returns most frequent english words
// Currently data is taken from open bases
func getStopWords() ([]string, error) {
	resp, err := http.Get("https://gist.githubusercontent.com/rg089/35e00abf8941d72d419224cfd5b5925d/raw/12d899b70156fd0041fa9778d657330b024b959c/stopwords.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func main() {
	flags := ParseFlags()

	err := flags.ValidateFlags()
	if err != nil {
		fmt.Printf("Failed to parse flags: %+v\n", err)
	}

	stemmer := SnowballStemmer{
		language:      English,
		stopWordsFunc: getStopWords,
	}
	stemmed, err := stemmer.Stem(strings.Split(flags.InitialString, " "))
	if err != nil {
		fmt.Printf("Failed to stem provided string: %+v\n", err)
	}

	fmt.Println(stemmed)
}
